package server

import (
	"io/fs"
	"net/http"

	"github.com/arunim-io/erp/internal/templates"
)

func RegisterRoutes(mux *http.ServeMux, staticFS fs.FS) error {
	fs, err := fs.Sub(staticFS, "static")
	if err != nil {
		return err
	}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServerFS(fs)))

	mux.HandleFunc("/", indexRoute)

	return nil
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	templates.RenderDefault(w, map[string]any{
		"PageTitle":  "ERP",
		"CurrentURL": r.URL.String(),
	})
}
