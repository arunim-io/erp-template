package server

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/arunim-io/erp-template/internal/config"
)

type discardWriter struct{}

func (d *discardWriter) Write(p []byte) (int, error) { return len(p), nil }

func newTestLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
}

func TestNewServerConfigApplied(t *testing.T) {
	ctx := context.Background()
	logger := newTestLogger()
	sm := scs.New()

	cfg := &config.ServerConfig{
		Host:              "127.0.0.1",
		Port:              9000,
		IdleTimeout:       time.Second,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      3 * time.Second,
		ReadHeaderTimeout: 4 * time.Second,
	}

	srv, err := New(ctx, logger, sm, nil, config.ModeDev, cfg)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	if srv == nil {
		t.Fatalf("expected non-nil server")
	}

	if srv.Addr != cfg.Addr() {
		t.Fatalf("expected Addr %q, got %q", cfg.Addr(), srv.Addr)
	}

	if srv.Handler == nil {
		t.Fatalf("expected non-nil handler")
	}

	if srv.IdleTimeout != cfg.IdleTimeout {
		t.Fatalf("expected IdleTimeout %s, got %s", cfg.IdleTimeout, srv.IdleTimeout)
	}

	if srv.ReadTimeout != cfg.ReadTimeout {
		t.Fatalf("expected ReadTimeout %s, got %s", cfg.ReadTimeout, srv.ReadTimeout)
	}

	if srv.WriteTimeout != cfg.WriteTimeout {
		t.Fatalf("expected WriteTimeout %s, got %s", cfg.WriteTimeout, srv.WriteTimeout)
	}

	if srv.ReadHeaderTimeout != cfg.ReadHeaderTimeout {
		t.Fatalf("expected ReadHeaderTimeout %s, got %s", cfg.ReadHeaderTimeout, srv.ReadHeaderTimeout)
	}
}
