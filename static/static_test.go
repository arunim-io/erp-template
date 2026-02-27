package static

import (
	"context"
	"log/slog"
	"testing"

	"github.com/arunim-io/erp-template/internal/config"
)

func TestRootDevModeUsesLocalDir(t *testing.T) {
	ctx := context.Background()
	mode := config.ModeDev
	logger := slog.New(slog.NewTextHandler(&discardWriter{}, nil))

	fs, err := Root(ctx, mode, logger)
	if err != nil {
		t.Fatalf("Root returned error in dev mode: %v", err)
	}

	if fs == nil {
		t.Fatalf("expected non-nil filesystem in dev mode")
	}
}

func TestRootProdModeUsesEmbeddedFS(t *testing.T) {
	ctx := context.Background()
	mode := config.ModeProd
	logger := slog.New(slog.NewTextHandler(&discardWriter{}, nil))

	fs, err := Root(ctx, mode, logger)
	if err != nil {
		t.Fatalf("Root returned error in prod mode: %v", err)
	}

	if fs == nil {
		t.Fatalf("expected non-nil filesystem in prod mode")
	}

	// Ensure we can open a known static file from the embedded filesystem.
	f, err := fs.Open("css/main.css")
	if err != nil {
		t.Fatalf("expected to open embedded css/main.css, got error: %v", err)
	}
	_ = f.Close()
}

// discardWriter implements io.Writer and discards all bytes.
type discardWriter struct{}

func (d *discardWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
