package erp

import (
	"net/http"

	"github.com/arunim-io/erp/internal/templates/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RootRouter(app *App) *chi.Mux {
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
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static/dist")),
		),
	)

	r.Get("/", IndexRoute(app))

	return r
}

func IndexRoute(app *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		users, _ := app.DB.Queries.ListUsers(ctx)

		pages.Index(len(users)).Render(ctx, w)
	}
}
