package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var greet http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { //nolint:revive // Just an useless warning
	_, _ = fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	const timeout = 5 * time.Second

	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	mux := http.NewServeMux()

	mux.HandleFunc("/", greet)

	addr := net.JoinHostPort("localhost", "8000")
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: timeout,
	}

	go func() {
		log.Info("Started server...", "addr", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Couldn't listen: %v\n", "error", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", "error", err)
	}

	log.Info("Server has successfully shut down.")
}
