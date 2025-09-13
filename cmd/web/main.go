package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	erp "github.com/arunim-io/erp/internal/core"
)

func main() {
	app := erp.NewApp()
	logger := app.Logger

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	s := http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", "8000"),
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
