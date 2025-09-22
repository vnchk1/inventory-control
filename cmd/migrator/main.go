package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	"github.com/vnchk1/inventory-control/internal/config"
	"github.com/vnchk1/inventory-control/internal/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(models.ErrEnvLoad)
	}

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatalf("main: %v", models.ErrCfgPath)
	}

	err = godotenv.Load(cfgPath)
	if err != nil {
		log.Println(models.ErrEnvLoad)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading migrator config: %v\n", err)
	}

	if err = runMigrations(cfg); err != nil {
		log.Fatalf("Error running migration: %v\n", err)
	}
}

func runMigrations(cfg *config.Config) (err error) {
	connStr := config.ConnStr(cfg)

	connConfig, err := pgx.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("migrator: error parsing config: %w", err)
	}

	db := stdlib.OpenDB(*connConfig)
	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error setting postgres dialect: %w", err)
	}

	if err = goose.Up(db, cfg.Migrator.Path); err != nil {
		return fmt.Errorf("migrator: UP error: %w", err)
	}

	return nil
}
