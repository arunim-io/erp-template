package app

import (
	"log/slog"

	"aidanwoods.dev/go-paseto"
)

type App struct {
	Logger   *slog.Logger
	Settings *Settings
	DB       *DB
	Key      *paseto.V4SymmetricKey
}

func NewApp(logger *slog.Logger) (*App, error) {
	app := &App{Logger: logger}

	settings, err := Load()
	if err != nil {
		return nil, err
	}
	app.Settings = settings

	db, err := GetDB(settings.Database.URI)
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

	return app, nil
}
