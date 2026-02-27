package logging

import (
	"log/slog"
	"testing"
)

func TestNewLoggerReturnsLogger(t *testing.T) {
	logger := NewLogger(slog.LevelInfo, false)
	if logger == nil {
		t.Fatalf("expected non-nil logger")
	}

	loggerConcise := NewLogger(slog.LevelDebug, true)
	if loggerConcise == nil {
		t.Fatalf("expected non-nil concise logger")
	}
}

func TestSchemaConciseFlag(t *testing.T) {
	s1 := Schema(false)
	s2 := Schema(true)

	if s1 == nil || s2 == nil {
		t.Fatalf("expected non-nil schemas")
	}

	if s1 == s2 {
		t.Fatalf("expected different schema pointers when concise flag differs")
	}
}
