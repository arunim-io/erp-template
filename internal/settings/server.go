package settings

import (
	"net"
	"strconv"
)

const defaultPort int16 = 8000

func defaultServerConfig() ServerConfig {
	return ServerConfig{
		Host: "0.0.0.0",
		Port: defaultPort,
	}
}

// ServerConfig contains the configuration of the server.
type ServerConfig struct {
	// The host at which the server will run.
	Host string `koanf:"host"`
	// The port at which the server will run.
	Port int16 `koanf:"port"`
}

// ServerAddress returns the IP address of the server.
func (cfg ServerConfig) ServerAddress() string {
	return net.JoinHostPort(cfg.Host, strconv.Itoa(int(cfg.Port)))
}
