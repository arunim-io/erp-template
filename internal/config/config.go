package config

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

type Config struct {
	Server  ServerConfig
	Logging LoggingConfig
}

func Default() Config {
	const (
		port              = 8000
		idleTimeout       = time.Minute
		readTimeout       = 10 * time.Second
		writeTimeout      = 30 * time.Second
		readHeaderTimeout = 2 * time.Second
	)

	return Config{
		Server: ServerConfig{
			Host:              "localhost",
			Port:              port,
			IdleTimeout:       idleTimeout,
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
			ReadHeaderTimeout: readHeaderTimeout,
		},
		Logging: LoggingConfig{
			Level: slog.LevelDebug,
		},
	}
}

func Load() (*Config, error) {
	var cfg Config

	_ = k.Load(structs.Provider(Default(), "koanf"), nil)

	_ = k.Load(env.Provider(".", env.Opt{
		Prefix: "ERP_",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(
				strings.TrimPrefix(k, "MYVAR_")), "_", ".")

			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)

	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}

	return &cfg, nil
}

type ServerConfig struct {
	Host              string
	Port              int16
	IdleTimeout       time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
}

func (c *ServerConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port)))
}

type LoggingConfig struct {
	Level slog.Leveler
}
