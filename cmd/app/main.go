package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vnchk1/inventory-control/config"
	"github.com/vnchk1/inventory-control/internal/db"
	logging "github.com/vnchk1/inventory-control/internal/logger"
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

	logger := logging.NewLogger(cfg.LogLevel)

	conn, err := db.InitDB(db.ConnStr(cfg))
	if err != nil {
		log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer conn.Close(context.Background())

	exampleProduct := &models.Product{
		ID:         17,
		Name:       "guesttttt",
		Quantity:   44,
		CategoryID: 1,
	}

	//err = repo.Create(conn, logger, exampleProduct)
	//if err != nil {
	//	log.Fatalf("Unable to create product: %v\n", err)
	//}

	exampleId := 17

	err = repo.Update(conn, logger, exampleProduct)
	if err != nil {
		log.Fatalf("Error updating product: %v\n", err)
	}

	err = repo.Read(conn, logger, exampleId)
	if err != nil {
		log.Fatalf("Unable to read product: %v\n", err)
	}

}
