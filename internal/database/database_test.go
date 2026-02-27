package database

import (
	"context"
	"strings"
	"testing"

	"github.com/arunim-io/erp-template/internal/config"
)

func TestNewRejectsInvalidDSN(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	_, err := New(ctx, " not-a-valid-dsn ", config.ModeDev)
	if err == nil {
		t.Fatalf("expected error for invalid DSN, got nil")
	}

	if !strings.Contains(err.Error(), "unable to parse database url") {
		t.Fatalf("expected parse error, got %v", err)
	}
}

func TestNewRequiresTLSInProd(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// sslmode=disable should result in nil TLSConfig, which is rejected in production mode.
	_, err := New(ctx, "postgres://user:pass@localhost:5432/dbname?sslmode=disable", config.ModeProd)
	if err == nil {
		t.Fatalf("expected error when TLS is not configured in production mode, got nil")
	}
}
