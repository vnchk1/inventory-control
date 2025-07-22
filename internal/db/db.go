package db

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

func connStr(cfg *config.Config) string {
	return fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}

func NewDB(cfg *config.Config) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connStr(cfg))
	if err != nil {
		return nil, fmt.Errorf("db connect error: %v\n", err)
	}

	return &DB{pool: pool}, nil
}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.pool.QueryRow(ctx, query, args...)
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, query, args...)
}

func (d *DB) Close(ctx context.Context) {
	d.pool.Close()
}
