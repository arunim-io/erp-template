package core

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/arunim-io/erp-template/internal/database"
)

func Router(queries *database.Queries) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", templ.Handler(HomePage()).ServeHTTP)

	return r
}
