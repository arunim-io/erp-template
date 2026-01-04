package server

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arunim-io/erp/internal/templates"
)

func Router(staticFS fs.FS) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/login", loginRoute)

	fs, err := fs.Sub(staticFS, "static")
	if err != nil {
		return nil, err
	}

	r.Handle("/static/*", http.StripPrefix(
		"/static/",
		http.FileServerFS(fs),
	))

	return r, nil
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "index.html", map[string]any{
		"PageTitle":  "ERP",
		"CurrentURL": r.URL.String(),
	})
}

func loginRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "login.html", map[string]any{
		"PageTitle":  "ERP - Login",
		"CurrentURL": r.URL.String(),
	})
}
