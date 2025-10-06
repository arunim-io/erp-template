package auth

import (
	"net/http"

	"github.com/arunim-io/erp/internal/app"
	pages "github.com/arunim-io/erp/internal/auth/templates/pages"
	"github.com/go-chi/chi/v5"
)

func Router(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.HandleFunc("/login", LoginRoute(app))

	return r
}

func LoginRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pages.LoginPage().Render(r.Context(), w)
		case http.MethodPost:
		}
	}
}
