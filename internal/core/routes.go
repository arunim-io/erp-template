package core

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", templ.Handler(HomePage()).ServeHTTP)

	return r
}
