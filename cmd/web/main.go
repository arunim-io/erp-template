package main

import (
	"log/slog"
	"os"

	"github.com/arunim-io/erp/internal/app"
	"github.com/arunim-io/erp/internal/server"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app, err := app.New(logger)
	if err != nil {
		logger.Error("Error while initializing...", "error", err)
		os.Exit(1)
	}

	server.New(app).Run()
}
