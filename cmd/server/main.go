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

	"github.com/arunim-io/erp-template/internal/config"
	"github.com/arunim-io/erp-template/internal/database"
	"github.com/arunim-io/erp-template/internal/logging"
	"github.com/arunim-io/erp-template/internal/server"
	"github.com/arunim-io/erp-template/internal/session"
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

	cfg, err := config.Load(slog.Default())
	if err != nil {
		return err
	}
	mode := cfg.Mode

	logger := logging.NewLogger(cfg.Logging.Level, !mode.IsDev()).
		With(slog.String("env", mode.String()))

	logger.DebugContext(ctx, "Config loaded", "data", cfg)

	db, err := database.New(ctx, cfg.Database.URL, cfg.Mode)
	if err != nil {
		return err
	}

	logger.DebugContext(ctx, "Database connected")

	sm := session.New(db, !cfg.Mode.IsProd(), cfg.SessionCookie)

	svr, err := server.New(ctx, logger, sm, db.Queries, cfg.Mode, cfg.Server)
	if err != nil {
		return err
	}

	go func() {
		logger.InfoContext(ctx, "Server running", "address", svr.Addr)
		if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error while listening: %v", err)
		}
	}()

	<-ctx.Done()
	logger.InfoContext(ctx, "Shutting down server")
	stop()

	const timeout = 10 * time.Second
	shutdownCtx, cancel := context.WithTimeout(rootCtx, timeout)
	defer cancel()

	if err := svr.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	logger.InfoContext(shutdownCtx, "Server sucessfully shut down")

	return nil
}
