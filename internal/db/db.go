package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/vnchk1/inventory-control/config"
)

func ConnStr(cfg *config.Config) string {
	return fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}

func InitDB(ConnStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), ConnStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return conn, nil
}
