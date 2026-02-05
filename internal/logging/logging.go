package logging

import (
	"log/slog"
	"os"

	"github.com/go-chi/httplog/v3"
)

func NewLogger(level slog.Leveler, concise bool) *slog.Logger {
	schema := Schema(concise)

	handlerOpts := &slog.HandlerOptions{
		AddSource:   concise,
		Level:       level,
		ReplaceAttr: schema.ReplaceAttr,
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, handlerOpts)

	if concise {
		handler = slog.NewJSONHandler(os.Stdout, handlerOpts)
	}

	return slog.New(handler)
}

func Schema(concise bool) *httplog.Schema {
	return httplog.SchemaECS.Concise(concise)
}
