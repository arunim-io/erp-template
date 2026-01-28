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
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(rootCtx context.Context) error {
	ctx, stop := signal.NotifyContext(rootCtx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	const timeout = 5 * time.Second
	server := &http.Server{
		Addr:              "localhost:8000",
		Handler:           http.HandlerFunc(handler),
		ReadHeaderTimeout: timeout,
	}

	go func() {
		logger.InfoContext(ctx, "Server running", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error while listening: %v", err)
		}
	}()

	<-ctx.Done()
	logger.InfoContext(ctx, "Shutting down server")
	stop()

	shutdownCtx, cancel := context.WithTimeout(rootCtx, timeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.ErrorContext(shutdownCtx, "Server forced to shutdown", "error", err)

		return err
	}

	logger.InfoContext(shutdownCtx, "Server sucessfully shut down")

	return nil
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
