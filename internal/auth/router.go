package auth

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"

	"github.com/arunim-io/erp-template/internal/auth/templates/pages"
)

func Router(
	formDecoder *form.Decoder,
	formValidator *validator.Validate,
	logger *slog.Logger,
	sm *scs.SessionManager,
) *chi.Mux {
	r := chi.NewRouter()

	r.Handle("/login", loginHandler(formDecoder, formValidator, logger))

	return r
}

func loginHandler(
	formDecoder *form.Decoder,
	formValidator *validator.Validate,
	logger *slog.Logger,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		switch r.Method {
		case http.MethodGet:
			_ = pages.Login(nil).Render(ctx, w)
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form data", http.StatusForbidden)
			}

			var form struct {
				Username string `validate:"required"`
				Password string `validate:"required,min=10,required" json:"-"`
			}

			if err := formDecoder.Decode(&form, r.Form); err != nil {
				http.Error(w, "Unable to parse form data", http.StatusNotAcceptable)
			}

			if err := formValidator.StructCtx(ctx, form); err != nil {
				var errs validator.ValidationErrors

				if errors.As(err, &errs) {
					_ = pages.Login(errs).Render(ctx, w)
				}
			}

			logger.DebugContext(ctx, "login form successfully submitted", "data", form)

			_ = pages.Login(nil).Render(ctx, w)
		}
	}
}
