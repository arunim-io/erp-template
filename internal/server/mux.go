package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/arunim-io/erp-template/internal/auth"
	"github.com/arunim-io/erp-template/internal/database"
	"github.com/arunim-io/erp-template/internal/templates/pages"
)

func Mux(
	_ *database.Queries,
	staticRoot http.FileSystem,
	middlewares ...func(http.Handler) http.Handler,
) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(middlewares...)

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(staticRoot)))
	mux.Get("/", templ.Handler(pages.Home()).ServeHTTP)

	mux.Mount("/auth/", auth.Router())

	return mux
}
