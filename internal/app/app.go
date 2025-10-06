package app

import (
	"log/slog"

	"aidanwoods.dev/go-paseto"
	"github.com/alexedwards/scs/v2"
	"github.com/arunim-io/erp/internal/db"
	"github.com/arunim-io/erp/internal/settings"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
)

type App struct {
	Logger         *slog.Logger
	Settings       *settings.Settings
	DB             *db.DB
	Key            *paseto.V4SymmetricKey
	SessionManager *scs.SessionManager
	Form           struct {
		Decoder   *form.Decoder
		Validator *validator.Validate
	}
}

func New(logger *slog.Logger) (*App, error) {
	app := &App{Logger: logger}

	settings, err := settings.Load()
	if err != nil {
		return nil, err
	}
	app.Settings = settings

	db, err := db.New(settings.Database.URI)
	if err != nil {
		return nil, err
	}
	app.DB = db

	var secretKey paseto.V4SymmetricKey

	if settings.SecretKey == "" {
		secretKey = paseto.NewV4SymmetricKey()

		logger.Warn("Using a temporary Secret Key. Please set the secret key as soon as possible.", "tempSecretKey", secretKey.ExportHex())
	} else {
		key, err := paseto.V4SymmetricKeyFromHex(settings.SecretKey)
		if err != nil {
			logger.Error("This secret key is invalid...", "error", err)
			return nil, err
		}

		secretKey = key
	}

	app.Key = &secretKey

	sm := scs.New()

	sm.Cookie = settings.SessionCookie

	app.SessionManager = sm

	app.Form.Decoder = form.NewDecoder()
	app.Form.Validator = validator.New(validator.WithRequiredStructEnabled())

	return app, nil
}
