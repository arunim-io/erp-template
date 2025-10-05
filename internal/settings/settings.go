package settings

import (
	"os"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Loads the settings from the default values, from the environment variables & from a `.env` file if present.
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

// This struct contains the settings used throughout the server.
type Settings struct {
	// The root directory for the project. Useful for accessing files in the project.
	RootDir string `koanf:"root_dir"`
	// A secret key for this server. This is mainly used for security purposes.
	SecretKey     string            `koanf:"secret_key"`
	Server        ServerConfig      `koanf:"server"`
	Database      *DBConfig         `koanf:"db"`
	SessionCookie scs.SessionCookie `koanf:"session"`
}

func getDefault() (*Settings, error) {
	s := &Settings{
		Server:        DefaultServerConfig(),
		SessionCookie: scs.SessionCookie{},
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	s.RootDir = cwd

	db, err := DefaultDBConfig(cwd)
	if err != nil {
		return nil, err
	}
	s.Database = db

	return s, nil
}

func loadFromEnv(s *Settings) error {
	k := koanf.New(".")

	loadDotenv := func(path string) error {
		if _, err := os.Stat(path); err == nil {
			if err := k.Load(file.Provider(path), dotenv.Parser()); err != nil {
				return err
			}
		}

		return nil
	}

	if paths, err := filepath.Glob(".env.*"); err == nil && len(paths) > 0 {
		for _, path := range paths {
			if path != ".env.example" {
				loadDotenv(path)
			}
		}
	} else {
		loadDotenv(".env")
	}

	// Load from Env variables
	if err := k.Load(
		env.Provider("ERP_", ".", func(v string) string { return v[4:] }),
		nil,
	); err != nil {
		return err
	}

	if err := k.Unmarshal("", s); err != nil {
		return err
	}

	return nil
}
