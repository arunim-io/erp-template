package erp

import (
	"log/slog"
	"os"
)

type App struct {
	Logger *slog.Logger
}

func NewApp() *App {
	return &App{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}
