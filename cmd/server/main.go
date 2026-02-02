package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/arunim-io/erp-template/internal/config"
	"github.com/arunim-io/erp-template/internal/core"
	"github.com/arunim-io/erp-template/internal/database"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(rootCtx context.Context) error {
	ctx, stop := signal.NotifyContext(rootCtx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	schema := httplog.SchemaECS.Concise(true)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       cfg.Logging.Level,
		ReplaceAttr: schema.ReplaceAttr,
	}))
	slog.SetDefault(logger)

	logger.DebugContext(ctx, "Config loaded", "data", cfg)

	dbPoolCfg, err := pgxpool.ParseConfig(cfg.Database.URL)
	if err != nil {
		return fmt.Errorf("unable to parse database url: %w", err)
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, dbPoolCfg)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer dbPool.Close()
	if err := dbPool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	queries := database.New(dbPool)

	logger.DebugContext(ctx, "Database connected")

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(dbPool)

	mux := chi.NewMux()

	mux.Use(
		middleware.CleanPath,
		middleware.GetHead,
		middleware.StripSlashes,
		httplog.RequestLogger(logger, &httplog.Options{
			Schema:        schema,
			RecoverPanics: true,
		}),
		sessionManager.LoadAndSave,
	)

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Mount("/", core.Router(queries))

	server := &http.Server{
		Addr:              cfg.Server.Addr(),
		Handler:           mux,
		IdleTimeout:       cfg.Server.IdleTimeout,
		ReadTimeout:       cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
	}

	go func() {
		logger.InfoContext(ctx, "Server running", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error while listening: %v", err)
		}
	}()

	<-ctx.Done()
	logger.InfoContext(ctx, "Shutting down server")
	stop()

	const timeout = 10 * time.Second
	shutdownCtx, cancel := context.WithTimeout(rootCtx, timeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	logger.InfoContext(shutdownCtx, "Server sucessfully shut down")

	return nil
}
