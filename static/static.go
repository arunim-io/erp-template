package static

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/arunim-io/erp-template/internal/config"
)

//go:embed *
var staticFiles embed.FS

func Root(ctx context.Context, mode config.Mode, logger *slog.Logger) (http.FileSystem, error) {
	if mode.IsDev() {
		logger.DebugContext(ctx, "Loading static files from local filesystem")

		return http.Dir("static"), nil
	}

	logger.InfoContext(ctx, "loading static file from embedded filesystem")

	subFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		return nil, fmt.Errorf("unable to load embedded static files: %w", err)
	}

	return http.FS(subFS), nil
}
