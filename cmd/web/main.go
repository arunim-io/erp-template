package main

import (
	"log/slog"
	"os"

	"github.com/arunim-io/erp/internal/app"
	"github.com/arunim-io/erp/internal/server"
	"github.com/go-chi/httplog/v3"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelDebug,
		ReplaceAttr: httplog.SchemaECS.ReplaceAttr,
	}))

	app, err := app.New(logger)
	if err != nil {
		logger.Error("Error while initializing...", "error", err)
		os.Exit(1)
	}

	server.New(app).Run()
}
