package main

import (
	"context"
	"embed"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//go:embed templates/*.html
var templatesFS embed.FS

var templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))

var routes = map[string]http.HandlerFunc{
	"/": func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", map[string]any{"PageTitle": "ERP", "CurrentURL": r.URL.String()}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	},
}

func main() {
	const timeout = 5 * time.Second

	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	mux := http.NewServeMux()

	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
	}

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
