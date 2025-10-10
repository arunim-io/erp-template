package auth

import (
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/arunim-io/erp/internal/app"
	pages "github.com/arunim-io/erp/internal/auth/templates/pages"
	"github.com/go-chi/chi/v5"
)

func Router(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.HandleFunc("/login", LoginRoute(app))
	r.HandleFunc("/register", RegisterRoute(app))

	return r
}

type LoginForm struct {
	Method   string `form:"login-method" validate:"required"`
	Email    string `form:"email" validate:"email,omitempty"`
	Username string `form:"username" validate:"omitempty"`
	Password string `form:"password" validate:"required"`
}

func LoginRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pages.LoginPage().Render(r.Context(), w)
		case http.MethodPost:
			r.ParseForm()

			var data LoginForm
			if err := app.Form.Decoder.Decode(&data, r.PostForm); err != nil {
				http.Error(w, "Invalid form", http.StatusNotAcceptable)
			}

			app.Logger.Debug("login successful", "data", data)

			htmx.NewResponse().Refresh(true).Write(w)
		}
	}
}

type RegisterForm struct {
	Email    string `form:"email" validate:"email"`
	Username string `form:"username"`
	Password string `form:"password" validate:"required"`
}

func RegisterRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pages.RegisterPage().Render(r.Context(), w)
		case http.MethodPost:
			r.ParseForm()

			var data RegisterForm
			if err := app.Form.Decoder.Decode(&data, r.PostForm); err != nil {
				http.Error(w, "Invalid form", http.StatusNotAcceptable)
			}

			app.Logger.Debug("registration successful", "data", data)
		}
	}
}
