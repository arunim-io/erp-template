package main

import (
	"context"
	"embed"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arunim-io/erp/internal/templates"
)

//go:embed templates/**/*.html
var templatesFS embed.FS

var routes = map[string]http.HandlerFunc{
	"/": func(w http.ResponseWriter, r *http.Request) {
		templates.RenderDefault(w, map[string]any{
			"PageTitle":  "ERP",
			"CurrentURL": r.URL.String(),
		})
	},
}

func main() {
	templates.Init(templatesFS)

	const timeout = 5 * time.Second

	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	mux := http.NewServeMux()

	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
	}

	addr := net.JoinHostPort("localhost", "8080")
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
