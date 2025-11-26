package server

import (
	"net/http"

	"github.com/arunim-io/erp/internal/templates"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", indexRoute)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	templates.RenderDefault(w, map[string]any{
		"PageTitle":  "ERP",
		"CurrentURL": r.URL.String(),
	})
}
