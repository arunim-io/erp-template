package server

import (
	"net/http"

	"filippo.io/csrf"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"

	"github.com/arunim-io/erp/internal/app"
	"github.com/arunim-io/erp/internal/auth"
	"github.com/arunim-io/erp/internal/orm"
	"github.com/arunim-io/erp/internal/templates/pages"
)

func RootRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	r.Use(
		middleware.SupressNotFound(r),
		httplog.RequestLogger(app.Logger, &httplog.Options{
			Level:         app.Settings.LogLevel(),
			Schema:        httplog.SchemaECS,
			RecoverPanics: true,
		}),
		middleware.CleanPath,
		middleware.GetHead,
		middleware.Recoverer,
		middleware.RedirectSlashes,
		middleware.StripSlashes,
		app.SessionManager.LoadAndSave,
		func(h http.Handler) http.Handler {
			if app.Settings.Debug {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Cache-Control", "no-store")
					h.ServeHTTP(w, r)
				})
			}

			return h
		},
	)

	r.Handle(
		"/static/*",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static/dist")),
		),
	)

	r.Get("/", IndexRoute(app))

	r.Mount("/", auth.Router(app))

	return csrf.New().Handler(r)
}

func IndexRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		users, _ := app.DB.Queries.ListUsers(ctx, orm.ListUsersParams{})

		_ = pages.Index(len(users)).Render(ctx, w)
	}
}
