package erp

import "log/slog"

type App struct {
	Logger   *slog.Logger
	Settings *Settings
	DB       *DB
}

func NewApp(logger *slog.Logger) (*App, error) {
	settings, err := GetSettings()
	if err != nil {
		return nil, err
	}

	db, err := GetDB(settings.Database.URI)
	if err != nil {
		return nil, err
	}

	return &App{
		Logger:   logger,
		Settings: settings,
		DB:       db,
	}, nil
}
