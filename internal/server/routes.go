package server

import (
	"io/fs"
	"net/http"

	"github.com/arunim-io/erp/internal/templates"
)

func RegisterRoutes(mux *http.ServeMux, staticFS fs.FS) error {
	mux.HandleFunc("/", indexRoute)
	mux.HandleFunc("/login", loginRoute)

	fs, err := fs.Sub(staticFS, "static")
	if err != nil {
		return err
	}

	mux.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServerFS(fs),
	))

	return nil
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
