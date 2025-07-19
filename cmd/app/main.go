package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vnchk1/inventory-control/config"
	"github.com/vnchk1/inventory-control/internal/db"
	"github.com/vnchk1/inventory-control/internal/logger"
	"github.com/vnchk1/inventory-control/internal/models"
	"github.com/vnchk1/inventory-control/internal/repo/cruds"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	logger := logger.NewLogger(cfg.LogLevel)

	conn, err := db.InitDB(db.ConnStr(cfg))
	if err != nil {
		log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer conn.Close(context.Background())

	user := &models.Products{
		Id:         1,
		Name:       "guest",
		Price:      1000,
		Quantity:   1,
		CategoryId: 1,
	}

	err = repo.Create(conn, user, logger)
	if err != nil {
		log.Fatalf("Unable to create product: %v\n", err)
	}
}
