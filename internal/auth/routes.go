package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arunim-io/erp/internal/templates"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.HandleFunc("/login", loginRoute)

	return r
}

func loginRoute(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "login.html", map[string]any{"PageTitle": "ERP - Login"})
}
