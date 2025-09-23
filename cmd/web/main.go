package main

import (
	"log/slog"
	"os"

	erp "github.com/arunim-io/erp/internal/core"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app, err := erp.NewApp(logger)
	if err != nil {
		logger.Error("Error while initializing...", "error", err)
		os.Exit(1)
	}

	app.RunServer()
}
