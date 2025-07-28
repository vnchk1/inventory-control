package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vnchk1/inventory-control/internal/config"
	"log"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(cfg *config.Config) (*DB, error) {
	connStr := config.ConnStr(cfg)
	log.Println("Connecting to DB", connStr)
	pool, err := pgxpool.New(context.Background(), connStr)
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

func (d *DB) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	return d.pool.BeginTx(ctx, opts)
}

func (d *DB) Close(ctx context.Context) {
	d.pool.Close()
}
