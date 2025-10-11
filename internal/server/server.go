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

const (
	readTimeout  time.Duration = 30 * time.Second
	writeTimeout time.Duration = 30 * time.Second
	idleTimeout  time.Duration = 120 * time.Second
)

type Server struct {
	instance *http.Server
	app      *app.App
}

func New(app *app.App) *Server {
	return &Server{
		app: app,
		instance: &http.Server{
			Addr:         app.Settings.Server.ServerAddress(),
			Handler:      RootRouter(app),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
	}
}

func (s *Server) Run() {
	go func() {
		s.app.Logger.Info("Starting server", "address", s.instance.Addr)

		if err := s.instance.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.app.Logger.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan

	ctx, cancel := context.WithTimeout(context.Background(), app.CancelTimeout)
	defer cancel()

	if err := s.instance.Shutdown(ctx); err != nil {
		s.app.Logger.Error("Server forced to shut down", "error", err)
	}
}
