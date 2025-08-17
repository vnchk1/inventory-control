package main

import (
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/vnchk1/inventory-control/internal/config"
)

func main() {
	_ = godotenv.Load()

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatalf("CONFIG_PATH is required")
	}

	_ = godotenv.Load(cfgPath)

	cfg, err := config.LoadMigratorConfig()
	if err != nil {
		log.Fatalf("Error loading migrator config %v", err)
	}

	connStr := config.MigratorConnStr(cfg)

	connConfig, err := pgx.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Migrator: error parsing config: %v", err)
	}

	db := stdlib.OpenDB(*connConfig)
	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Error setting postgres dialect: %v", err)
	}

	if err = goose.Up(db, cfg.MigrationsPath); err != nil {
		log.Fatalf("Migrator: UP error: %v", err)
	}
}
