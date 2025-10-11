package settings

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Load Loads the settings from the default values, from the environment variables &
// from a `.env` file if present.
func Load() (*Settings, error) {
	s, err := getDefault()
	if err != nil {
		return nil, err
	}

	if err := loadFromEnv(s); err != nil {
		return nil, err
	}

	return s, nil
}

// Settings contains the settings used throughout the server.
type Settings struct {
	// A flag that controls parts of the server.
	Debug bool `koanf:"debug"`
	// The root directory for the project. Useful for accessing files in the project.
	RootDir string `koanf:"root_dir"`
	// A secret key for this server. This is mainly used for security purposes.
	SecretKey     string            `koanf:"secret_key"`
	Server        ServerConfig      `koanf:"server"`
	Database      *DBConfig         `koanf:"db"`
	SessionCookie scs.SessionCookie `koanf:"session"`
}

// LogLevel returns the current slog.Level of the server.
func (s *Settings) LogLevel() slog.Level {
	if s.Debug {
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

func getDefault() (*Settings, error) {
	s := &Settings{
		Server: defaultServerConfig(),
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	s.RootDir = cwd

	db, err := defaultDBConfig(cwd)
	if err != nil {
		return nil, err
	}

	s.Database = db

	s.SessionCookie = scs.SessionCookie{Secure: !s.Debug}

	return s, nil
}

func loadFromEnv(s *Settings) (err error) {
	k := koanf.New(".")

	paths, err := filepath.Glob(".env.*")
	if err != nil {
		return fmt.Errorf("globbing dotenv files: %w", err)
	}

	paths = append(paths, ".env")

	for _, path := range paths {
		if path == ".env.example" {
			continue
		}

		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				return nil
			}

			return fmt.Errorf("checking %q: %w", path, err)
		}

		if err := k.Load(file.Provider(path), dotenv.Parser()); err != nil {
			return fmt.Errorf("loading %q: %w", path, err)
		}
	}

	const envPrefix string = "ERP_"
	envPrefixLen := len(envPrefix)

	if err := k.Load(
		env.Provider(envPrefix, ".", func(v string) string {
			if len(v) > envPrefixLen {
				return v[envPrefixLen:]
			}

			return v
		}),
		nil,
	); err != nil {
		return fmt.Errorf("loading environment variables: %w", err)
	}

	if err := k.Unmarshal("", s); err != nil {
		return fmt.Errorf("unmarshalling settings: %w", err)
	}

	return nil
}
