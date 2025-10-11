package settings

import (
	"net/url"
	"path"
	"time"
)

func defaultDBConfig(cwd string) (*DBConfig, error) {
	dsn, err := url.Parse(
		"file:" + path.Join(cwd, "db.sqlite3") + "?_journal_mode=WAL&_foreign_keys=on",
	)
	if err != nil {
		return nil, err
	}

	return &DBConfig{URI: dsn, MaxAge: 0}, nil
}

// DBConfig contains the configuration of the database used by the server.
type DBConfig struct {
	// The URL of the database
	URI *url.URL `koanf:"url"`
	// The lifetime of a database connection, in seconds
	MaxAge time.Duration `koanf:"max_age"`
}
