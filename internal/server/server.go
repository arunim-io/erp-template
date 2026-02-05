package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"

	"github.com/arunim-io/erp-template/internal/config"
	"github.com/arunim-io/erp-template/internal/database"
	"github.com/arunim-io/erp-template/internal/logging"
	"github.com/arunim-io/erp-template/static"
)

func New(
	ctx context.Context,
	logger *slog.Logger,
	sm *scs.SessionManager,
	queries *database.Queries,
	mode config.Mode,
	cfg *config.ServerConfig,
) (*http.Server, error) {
	staticRoot, err := static.Root(ctx, mode, logger)
	if err != nil {
		return nil, err
	}

	mux := Mux(
		queries,
		staticRoot,
		middleware.CleanPath,
		middleware.GetHead,
		middleware.StripSlashes,
		httplog.RequestLogger(logger, &httplog.Options{
			Schema:        logging.Schema(!mode.IsDev()),
			RecoverPanics: true,
		}),
		sm.LoadAndSave,
	)

	return &http.Server{
		Addr:              cfg.Addr(),
		Handler:           mux,
		IdleTimeout:       cfg.IdleTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}, nil
}
