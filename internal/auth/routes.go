package auth

import (
	"errors"
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/arunim-io/erp/internal/app"
	pages "github.com/arunim-io/erp/internal/auth/templates/pages"
)

func Router(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.HandleFunc("/login", LoginRoute(app))
	r.HandleFunc("/register", RegisterRoute(app))

	return r
}

func LoginRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			_ = pages.LoginPage().Render(r.Context(), w)
		case http.MethodPost:
			_ = r.ParseForm()

			var data LoginForm
			if err := app.Form.Decoder.Decode(&data, r.PostForm); err != nil {
				http.Error(w, "Invalid form", http.StatusNotAcceptable)
			}

			app.Logger.Debug("login successful", "data", data)

			_ = htmx.NewResponse().Refresh(true).Write(w)
		}
	}
}

func RegisterRoute(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var formData RegisterForm

		switch r.Method {
		case http.MethodGet:
			_ = pages.RegisterPage().Render(ctx, w)
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			if err := app.Form.Decoder.Decode(&formData, r.PostForm); err != nil {
				http.Error(w, "Invalid form", http.StatusNotAcceptable)
			}

			if err := app.Form.Validator.StructCtx(ctx, formData); err != nil {
				app.Logger.Error("Form data", "err", err)

				var vErrs validator.ValidationErrors

				if ok := errors.As(err, &vErrs); ok {
					for _, fe := range vErrs {
						app.Logger.Error("Field error", "field", fe.Field(), "error", fe.Error())
					}
				}

				_ = pages.RegisterPage().Render(ctx, w)
			}
		}
	}
}
