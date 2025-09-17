package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	erp "github.com/arunim-io/erp/internal/core"
	templates "github.com/arunim-io/erp/templates/pages"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app, err := erp.NewApp(logger)
	if err != nil {
		logger.Error("Error while initializing...", "error", err)
	}
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/dist"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		users, _ := app.DB.Queries.ListUsers(ctx)

		templates.Index(len(users)).Render(ctx, w)
	})

	s := http.Server{
		Addr:         app.Settings.GetServerAddress(),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Info("Starting server", "address", s.Addr)

		if err := s.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			logger.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	quit_chan := make(chan os.Signal, 1)
	signal.Notify(quit_chan, syscall.SIGINT, syscall.SIGTERM)
	<-quit_chan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shut down", "error", err)
		os.Exit(1)
	}

	logger.Info("Server has exited...")
}
