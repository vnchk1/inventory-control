package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	inventoryApp "github.com/vnchk1/inventory-control/internal/app"
	"github.com/vnchk1/inventory-control/internal/config"
	logging "github.com/vnchk1/inventory-control/internal/logger"
	"github.com/vnchk1/inventory-control/internal/models"
)

func main() {
	app, err := Setup()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}

	inventoryApp.Shutdown(app)
}

func Setup() (app *inventoryApp.App, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Println(models.ErrEnvLoad)
	}

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		return nil, models.ErrCfgPath
	}

	err = godotenv.Load(cfgPath)
	if err != nil {
		log.Println(models.ErrEnvLoad)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	logger := logging.NewLogger(cfg.LogLevel)

	app, err = inventoryApp.NewApp(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("error creating app %w", err)
	}

	return app, nil
}
