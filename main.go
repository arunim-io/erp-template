package main

import (
	"context"
	"embed"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/arunim-io/erp/internal/server"
	"github.com/arunim-io/erp/internal/templates"
)

//go:embed templates/**/*.html
var templatesFS embed.FS

func main() {
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	if err := templates.Init(templatesFS); err != nil {
		log.Error("Failed to initialize templates", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	server.RegisterRoutes(mux)

	s := server.New(mux)

	go func() {
		log.Info("Started server...", "addr", s.Addr)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Server errored while starting", "error", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Error("Server errored during shutdown", "error", err)
	}

	log.Info("Server has successfully shut down.")
}
