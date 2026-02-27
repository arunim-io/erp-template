package session

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/arunim-io/erp-template/internal/config"
	"github.com/arunim-io/erp-template/internal/database"
)

func TestNewSessionManagerConfig(t *testing.T) {
	db := &database.DB{
		Pool: (*pgxpool.Pool)(nil),
	}

	lifetime := 42 * time.Minute
	cfg := &config.SessionCookieConfig{
		Lifetime: lifetime,
	}

	sm := New(db, true, cfg)
	if sm == nil {
		t.Fatalf("expected non-nil session manager")
	}

	if sm.Lifetime != lifetime {
		t.Fatalf("expected lifetime %s, got %s", lifetime, sm.Lifetime)
	}

	if !sm.Cookie.Secure {
		t.Fatalf("expected cookie.Secure to be true when secure flag is true")
	}

	if !sm.Cookie.Partitioned {
		t.Fatalf("expected cookie.Partitioned to be true when secure flag is true")
	}
}
