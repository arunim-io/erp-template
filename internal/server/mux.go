package server

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"

	"github.com/arunim-io/erp-template/internal/auth"
	"github.com/arunim-io/erp-template/internal/database"
	"github.com/arunim-io/erp-template/internal/templates/pages"
)

func Mux(
	_ *database.Queries,
	staticRoot http.FileSystem,
	formDecoder *form.Decoder,
	formValidator *validator.Validate,
	logger *slog.Logger,
	sm *scs.SessionManager,
	middlewares ...func(http.Handler) http.Handler,
) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(middlewares...)

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(staticRoot)))
	mux.Get("/", templ.Handler(pages.Home()).ServeHTTP)

	mux.Mount("/auth/", auth.Router(formDecoder, formValidator, logger, sm))

	return mux
}
