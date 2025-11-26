package server

import (
	"net"
	"net/http"
	"time"
)

const (
	DefaultTimeout     = 5 * time.Second
	defaultIdleTimeout = 120 * time.Second
)

func New(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:         net.JoinHostPort("localhost", "8080"),
		Handler:      mux,
		ReadTimeout:  DefaultTimeout,
		WriteTimeout: DefaultTimeout,
		IdleTimeout:  defaultIdleTimeout,
	}
}
