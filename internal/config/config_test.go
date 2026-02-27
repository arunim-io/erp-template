package config

import (
	"log/slog"
	"testing"
	"time"
)

func TestDefaultConfigValues(t *testing.T) {
	cfg := Default()

	if cfg.Mode != ModeDev {
		t.Fatalf("expected default mode %q, got %q", ModeDev, cfg.Mode)
	}

	if cfg.Server == nil {
		t.Fatalf("expected non-nil server config")
	}

	if cfg.Logging == nil {
		t.Fatalf("expected non-nil logging config")
	}

	if cfg.SessionCookie == nil {
		t.Fatalf("expected non-nil session cookie config")
	}

	if got := cfg.Server.Port; got <= 0 {
		t.Fatalf("expected positive server port, got %d", got)
	}

	if cfg.SessionCookie.Lifetime <= 0 {
		t.Fatalf("expected positive session cookie lifetime, got %s", cfg.SessionCookie.Lifetime)
	}
}

func TestModeHelpersAndString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantMode    Mode
		wantIsDev   bool
		wantIsProd  bool
		wantString  string
		expectError bool
	}{
		{
			name:       "development full",
			input:      "development",
			wantMode:   ModeDev,
			wantIsDev:  true,
			wantString: "development",
		},
		{
			name:       "development short",
			input:      "dev",
			wantMode:   ModeDev,
			wantIsDev:  true,
			wantString: "development",
		},
		{
			name:       "production full",
			input:      "production",
			wantMode:   ModeProd,
			wantIsProd: true,
			wantString: "production",
		},
		{
			name:       "production short",
			input:      "prod",
			wantMode:   ModeProd,
			wantIsProd: true,
			wantString: "production",
		},
		{
			name:        "unknown",
			input:       "staging",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m Mode
			err := m.UnmarshalText([]byte(tt.input))

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.input, err)
			}

			if m != tt.wantMode {
				t.Fatalf("expected mode %q, got %q", tt.wantMode, m)
			}

			if got := m.IsDev(); got != tt.wantIsDev {
				t.Fatalf("IsDev for %q: expected %v, got %v", tt.input, tt.wantIsDev, got)
			}

			if got := m.IsProd(); got != tt.wantIsProd {
				t.Fatalf("IsProd for %q: expected %v, got %v", tt.input, tt.wantIsProd, got)
			}

			if s := m.String(); s != tt.wantString {
				t.Fatalf("String for %q: expected %q, got %q", tt.input, tt.wantString, s)
			}
		})
	}
}

func TestServerConfigAddr(t *testing.T) {
	cfg := &ServerConfig{
		Host: "127.0.0.1",
		Port: 8080,
	}

	if got, want := cfg.Addr(), "127.0.0.1:8080"; got != want {
		t.Fatalf("Addr() = %q, want %q", got, want)
	}
}

func TestLoggingConfigLogValue(t *testing.T) {
	lc := LoggingConfig{
		Level: slog.LevelDebug,
	}

	val := lc.LogValue()
	if val.Kind() != slog.KindGroup {
		t.Fatalf("expected group slog.Value, got %v", val.Kind())
	}
}

func TestSessionCookieConfigLogValue(t *testing.T) {
	sc := SessionCookieConfig{
		Lifetime: time.Minute,
	}

	val := sc.LogValue()
	if val.Kind() != slog.KindGroup {
		t.Fatalf("expected group slog.Value, got %v", val.Kind())
	}
}
