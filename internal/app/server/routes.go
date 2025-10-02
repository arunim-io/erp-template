package server

import (
	"net/http"

	"github.com/a-h/templ"
	erp "github.com/arunim-io/erp/internal/app"
	auth "github.com/arunim-io/erp/internal/auth/templates/pages"
	"github.com/arunim-io/erp/internal/orm"
	"github.com/arunim-io/erp/internal/templates/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RootRouter(app *erp.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.SupressNotFound(r),
		middleware.Logger,
		middleware.CleanPath,
		middleware.GetHead,
		middleware.Recoverer,
		middleware.RedirectSlashes,
		middleware.StripSlashes,
	)

	r.Handle(
		"/static/*",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static/dist")),
		),
	)

	r.Get("/", IndexRoute(app))
	r.Get("/login", templ.Handler(auth.LoginPage()).ServeHTTP)

	return r
}

func IndexRoute(app *erp.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		users, _ := app.DB.Queries.ListUsers(ctx, orm.ListUsersParams{})

		pages.Index(len(users)).Render(ctx, w)
	}
}
