package erp

import (
	"database/sql"
	"net/url"

	"github.com/arunim-io/erp/internal/orm"
	_ "modernc.org/sqlite"
)

type DB struct {
	db      *sql.DB
	Queries *orm.Queries
}

func GetDB(dsn *url.URL) (*DB, error) {
	db, err := sql.Open("sqlite", dsn.String())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	queries := orm.New(db)

	return &DB{db: db, Queries: queries}, nil
}
