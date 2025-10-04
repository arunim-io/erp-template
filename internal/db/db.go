package db

import (
	"database/sql"
	"net/url"

	"github.com/arunim-io/erp/internal/orm"
	_ "modernc.org/sqlite"
)

type DB struct {
	instance *sql.DB
	Queries  *orm.Queries
}

func New(dsn *url.URL) (*DB, error) {
	db, err := sql.Open("sqlite", dsn.String())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	queries := orm.New(db)

	return &DB{instance: db, Queries: queries}, nil
}
