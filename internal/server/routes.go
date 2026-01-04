package server

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arunim-io/erp/internal/templates"
)

func Router(staticFS fs.FS) (*chi.Mux, error) {
	r := chi.NewRouter()

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
	r.HandleFunc("/login", loginRoute)

	return r, nil
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "index.html", map[string]any{
		"PageTitle":  "ERP",
		"CurrentURL": r.URL,
	})
}

func loginRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "login.html", map[string]any{
		"PageTitle": "ERP - Login",
	})
}

func notFoundRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "404.html", map[string]any{
		"PageTitle": "ERP - Page not found",
	})
}
