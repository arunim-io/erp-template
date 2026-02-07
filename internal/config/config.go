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
	Mode          Mode                 `koanf:"mode"`
	Server        *ServerConfig        `koanf:"server"`
	Logging       *LoggingConfig       `koanf:"logging"`
	Database      *DBConfig            `koanf:"database"`
	SessionCookie *SessionCookieConfig `koanf:"session_cookie"`
}

func Default() Config {
	const (
		port                  = 8000
		idleTimeout           = time.Minute
		readTimeout           = 10 * time.Second
		writeTimeout          = 30 * time.Second
		readHeaderTimeout     = 2 * time.Second
		sessionCookieLifetime = 24 * time.Hour
	)

	return Config{
		Mode: ModeDev,
		Server: &ServerConfig{
			Host:              "localhost",
			Port:              port,
			IdleTimeout:       idleTimeout,
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
			ReadHeaderTimeout: readHeaderTimeout,
		},
		Logging: &LoggingConfig{
			Level: slog.LevelDebug,
		},
		SessionCookie: &SessionCookieConfig{
			Lifetime: sessionCookieLifetime,
		},
	}
}

func Load(logger *slog.Logger) (*Config, error) {
	var cfg Config

	_ = k.Load(structs.Provider(Default(), "koanf"), nil)

	_ = k.Load(env.Provider(".", env.Opt{
		Prefix: "ERP_",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(
				strings.TrimPrefix(k, "ERP_")), "_", ".")

			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)

	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}

	if cfg.Mode.IsProd() {
		if cfg.Logging.Level == slog.LevelDebug {
			logger.Warn("debug logs are not shown in production")
			cfg.Logging.Level = slog.LevelInfo
		}
	}

	return &cfg, nil
}

func (c Config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("server", c.Server),
		slog.Any("logging", c.Logging),
		slog.Any("database", c.Database),
		slog.Any("session_cookie", c.SessionCookie),
	)
}

type Mode string

const (
	ModeDev  Mode = "development"
	ModeProd Mode = "production"
)

func (m *Mode) IsDev() bool  { return *m == ModeDev }
func (m *Mode) IsProd() bool { return *m == ModeProd }

func (m *Mode) String() string {
	switch *m {
	case ModeDev:
		return "development"
	case ModeProd:
		return "production"
	default:
		return ""
	}
}

func (m *Mode) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "development", "dev":
		*m = ModeDev
	case "production", "prod":
		*m = ModeProd
	default:
		return fmt.Errorf("unknown mode :%s", text)
	}

	return nil
}

type ServerConfig struct {
	Host              string        `koanf:"host"`
	Port              int16         `koanf:"port"`
	IdleTimeout       time.Duration `koanf:"idle_timeout"`
	ReadTimeout       time.Duration `koanf:"read_timeout"`
	WriteTimeout      time.Duration `koanf:"write_timeout"`
	ReadHeaderTimeout time.Duration `koanf:"read_header_timeout"`
}

func (c *ServerConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port)))
}

func (c ServerConfig) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("host", c.Host),
		slog.Int("port", int(c.Port)),
		slog.Duration("idle_timeout", c.IdleTimeout),
		slog.Duration("read_timeout", c.ReadTimeout),
		slog.Duration("write_timeout", c.WriteTimeout),
		slog.Duration("read_header_timeout", c.ReadHeaderTimeout),
	)
}

type LoggingConfig struct {
	Level slog.Leveler `koanf:"level"`
}

func (c LoggingConfig) LogValue() slog.Value {
	return slog.GroupValue(slog.Any("level", c.Level))
}

type DBConfig struct {
	URL string `koanf:"url"`
}

func (c DBConfig) LogValue() slog.Value {
	return slog.GroupValue(slog.String("url", c.URL))
}

type SessionCookieConfig struct {
	Lifetime time.Duration `koanf:"lifetime"`
}

func (c SessionCookieConfig) LogValue() slog.Value {
	return slog.GroupValue(slog.Duration("lifetime", c.Lifetime))
}
