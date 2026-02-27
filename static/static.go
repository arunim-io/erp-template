package static

import (
	"context"
	"embed"
	"log/slog"
	"net/http"

	"github.com/arunim-io/erp-template/internal/config"
)

//go:embed css/*
var staticFiles embed.FS

func Root(ctx context.Context, mode config.Mode, logger *slog.Logger) (http.FileSystem, error) {
	if mode.IsDev() {
		logger.DebugContext(ctx, "Loading static files from local filesystem")

		return http.Dir("static"), nil
	}

	logger.InfoContext(ctx, "loading static file from embedded filesystem")

	return http.FS(staticFiles), nil
}
