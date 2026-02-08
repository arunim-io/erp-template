package auth

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/arunim-io/erp-template/internal/auth/templates/pages"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/login", templ.Handler(pages.Login()).ServeHTTP)

	return r
}
