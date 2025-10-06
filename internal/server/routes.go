package server

import (
	"log/slog"
	"net/http"

	"github.com/arunim-io/erp/internal/app"
	"github.com/arunim-io/erp/internal/auth"
	"github.com/arunim-io/erp/internal/orm"
	"github.com/arunim-io/erp/internal/templates/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
	"github.com/gorilla/csrf"
)

func RootRouter(app *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.SupressNotFound(r),
		httplog.RequestLogger(app.Logger, &httplog.Options{
			Level:         slog.LevelDebug,
			Schema:        httplog.SchemaECS,
			RecoverPanics: true,
		}),
		middleware.CleanPath,
		middleware.GetHead,
		middleware.Recoverer,
		middleware.RedirectSlashes,
		middleware.StripSlashes,
		app.SessionManager.LoadAndSave,
		csrf.Protect(
			app.Key.ExportBytes(),
			csrf.HttpOnly(false),
			csrf.Secure(false),
			csrf.SameSite(csrf.SameSiteLaxMode),
		),
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

	return r
}

func IndexRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		users, _ := app.DB.Queries.ListUsers(ctx, orm.ListUsersParams{})

		pages.Index(len(users)).Render(ctx, w)
	}
}
