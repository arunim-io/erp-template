package server

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

type nopMiddleware struct{}

func (nopMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler) {
	next.ServeHTTP(w, r)
}

func repoStaticDir(t *testing.T) string {
	t.Helper()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("unable to determine test file location")
	}

	// This file is at internal/server/mux_test.go; static/ is two levels up.
	return filepath.Clean(filepath.Join(filepath.Dir(thisFile), "../../static"))
}

func TestMuxServesHomePage(t *testing.T) {
	fs := http.Dir(repoStaticDir(t))
	mux := Mux(nil, fs)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if body := rr.Body.String(); body == "" {
		t.Fatalf("expected non-empty response body for home page")
	}
}

func TestMuxServesStaticFiles(t *testing.T) {
	fs := http.Dir(repoStaticDir(t))
	mux := Mux(nil, fs)

	req := httptest.NewRequest(http.MethodGet, "/static/css/main.css", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d for static file, got %d", http.StatusOK, rr.Code)
	}
}
