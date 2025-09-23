package app

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Settings struct {
	RootDir string `koanf:"root_dir"`

	Server struct {
		Host string `koanf:"host"`
		Port int    `koanf:"port"`
	} `koanf:"server"`

	Database struct {
		URI *url.URL `koanf:"url"`
	} `koanf:"db"`
}

func Load() (*Settings, error) {
	k := koanf.New(".")

	s := &Settings{}

	s.Server.Host = "0.0.0.0"
	s.Server.Port = 8000

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	s.RootDir = cwd

	dsn, err := url.Parse("file:" + path.Join(cwd, "db.sqlite3") + "?_journal_mode=WAL&_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	s.Database.URI = dsn

	// Load from dotenv
	envFile := ".env"
	if _, err := os.Stat(envFile); err == nil {
		if err := k.Load(file.Provider(envFile), dotenv.Parser()); err != nil {
			return nil, err
		}
	}

	// Load from Env variables
	if err := k.Load(
		env.Provider("ERP_", ".", func(s string) string { return s[4:] }),
		nil,
	); err != nil {
		return nil, err
	}

	if err := k.Unmarshal("", s); err != nil {
		return nil, err
	}

	return s, nil
}

func (s Settings) ServerAddress() string {
	return net.JoinHostPort(s.Server.Host, fmt.Sprint(s.Server.Port))
}
