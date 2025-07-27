package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vnchk1/inventory-control/internal/app"
	"github.com/vnchk1/inventory-control/internal/config"
	logging "github.com/vnchk1/inventory-control/internal/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	//логгер
	logger := logging.NewLogger(cfg.LogLevel)

	app := app.NewApp(cfg, logger)

	app.Run()

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Shutting down app...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.Stop(ctx)
}
