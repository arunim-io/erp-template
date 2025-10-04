package settings

import (
	"fmt"
	"net"
)

// Returns the default config for the server.
func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Host: "0.0.0.0",
		Port: 8000,
	}
}

// The configuration of the server.
type ServerConfig struct {
	// The host at which the server will run.
	Host string `koanf:"host"`
	// The port at which the server will run.
	Port int `koanf:"port"`
}

// Returns the IP address of the server.
func (cfg ServerConfig) ServerAddress() string {
	return net.JoinHostPort(cfg.Host, fmt.Sprint(cfg.Port))
}
