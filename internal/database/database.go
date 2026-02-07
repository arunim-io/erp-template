package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/arunim-io/erp-template/internal/config"
	"github.com/arunim-io/erp-template/internal/database/postgres"
)

type DB struct {
	*pgxpool.Pool

	Queries *postgres.Queries
}

func New(ctx context.Context, dsn string, mode config.Mode) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database url: %w", err)
	}

	if mode.IsProd() && cfg.ConnConfig.TLSConfig == nil {
		return nil, errors.New("sslmode must be configured in production")
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{Pool: pool, Queries: postgres.New(pool)}, nil
}
