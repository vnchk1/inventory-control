package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vnchk1/inventory-control/internal/config"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(cfg *config.Config) (*DB, error) {
	connStr := config.ConnStr(cfg)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	return &DB{pool: pool}, nil
}

//nolint:ireturn // возвращаем pgx.Row для доступа ко всем методам pgx
func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.pool.QueryRow(ctx, query, args...)
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, query, args...)
}

func (d *DB) GetConnString(cfg *config.Config) string {
	return config.ConnStr(cfg)
}

func (d *DB) Close() {
	d.pool.Close()
}
