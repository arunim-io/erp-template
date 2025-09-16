package erp

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path"
)

type Settings struct {
	Server struct {
		Host string
		Port int
	}
	Database struct {
		URI *url.URL
	}
}

func GetSettings() (*Settings, error) {
	s := &Settings{}

	s.Server.Host = "0.0.0.0"
	s.Server.Port = 8000

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dsn, err := url.Parse(fmt.Sprintf("file:%s?_journal_mode=WAL&_foreign_keys=on", path.Join(cwd, "db.sqlite3")))
	if err != nil {
		return nil, err
	}

	s.Database.URI = dsn

	return s, nil
}

func (s Settings) GetServerAddress() string {
	return net.JoinHostPort(s.Server.Host, fmt.Sprint(s.Server.Port))
}
