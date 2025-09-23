package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arunim-io/erp/internal/app"
)

type Server struct {
	instance *http.Server
	app      *app.App
}

func New(app *app.App) *Server {
	return &Server{
		instance: &http.Server{
			Addr:         app.Settings.ServerAddress(),
			Handler:      RootRouter(app),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		app: app,
	}
}

func (s *Server) Run() {
	go func() {
		s.app.Logger.Info("Starting server", "address", s.instance.Addr)

		if err := s.instance.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {

			s.app.Logger.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	quit_chan := make(chan os.Signal, 1)
	signal.Notify(quit_chan, syscall.SIGINT, syscall.SIGTERM)
	<-quit_chan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.instance.Shutdown(ctx); err != nil {
		s.app.Logger.Error("Server forced to shut down", "error", err)
		os.Exit(1)
	}

	s.app.Logger.Info("Server has exited...")
}
