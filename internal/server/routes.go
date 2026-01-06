package server

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/arunim-io/erp/internal/auth"
	"github.com/arunim-io/erp/internal/templates"
)

func Router(staticFS fs.FS) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.CleanPath, middleware.GetHead, middleware.Logger)

	fs, err := fs.Sub(staticFS, "static")
	if err != nil {
		return nil, err
	}

	r.Handle("/static/*", http.StripPrefix(
		"/static/",
		http.FileServerFS(fs),
	))

	r.NotFound(notFoundRoute)

	r.HandleFunc("/", indexRoute)

	r.Mount("/auth/", auth.Router())

	return r, nil
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "index.html", map[string]any{"PageTitle": "ERP"})
}

func notFoundRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "404.html", map[string]any{"PageTitle": "ERP - Page not found"})
}
