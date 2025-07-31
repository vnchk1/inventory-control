package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	apppkg "github.com/vnchk1/inventory-control/internal/app"
	"github.com/vnchk1/inventory-control/internal/config"
	logging "github.com/vnchk1/inventory-control/internal/logger"
)

const (
	ShutdownTimeoutSeconds = 5
)

func main() {
	_ = godotenv.Load()

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatalf("CONFIG_PATH is required")
	}

	_ = godotenv.Load(cfgPath)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	logger := logging.NewLogger(cfg.LogLevel)

	app, err := apppkg.NewApp(cfg, logger)
	if err != nil {
		log.Fatalf("Error creating app %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Shutting down app...")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeoutSeconds*time.Second)
	defer cancel()

	app.Stop(ctx)
}
