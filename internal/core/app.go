package erp

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arunim-io/erp/internal/templates/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Logger   *slog.Logger
	Settings *Settings
	DB       *DB
	Server   *http.Server
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

	app.Server = &http.Server{
		Addr:         app.Settings.ServerAddress(),
		Handler:      RootRouter(app),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return app, nil
}

func (app App) RunServer() {
	s := app.Server

	go func() {
		app.Logger.Info("Starting server", "address", s.Addr)

		if err := s.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			app.Logger.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	quit_chan := make(chan os.Signal, 1)
	signal.Notify(quit_chan, syscall.SIGINT, syscall.SIGTERM)
	<-quit_chan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		app.Logger.Error("Server forced to shut down", "error", err)
		os.Exit(1)
	}

	app.Logger.Info("Server has exited...")
}
