package app

import (
	"log/slog"
)

type App struct {
	Logger   *slog.Logger
	Settings *Settings
	DB       *DB
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

	return app, nil
}
