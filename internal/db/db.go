package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	// Load sqlite3 driver.
	_ "modernc.org/sqlite"

	"github.com/arunim-io/erp/internal/orm"
)

type DB struct {
	instance *sql.DB
	Queries  *orm.Queries
}

func New(ctx context.Context, dsn *url.URL) (db *DB, err error) {
	i, err := sql.Open("sqlite", dsn.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		if e := i.Close(); e != nil {
			err = fmt.Errorf("%w; also failed to close DB: %w", err, e)
		}
	}()

	if err = i.PingContext(ctx); err != nil {
		return nil, err
	}

	return &DB{instance: i, Queries: orm.New(i)}, nil
}
