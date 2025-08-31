package main

import (
	"context"
	"errors"
	"github.com/vnchk1/inventory-control/internal/models"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	inventoryApp "github.com/vnchk1/inventory-control/internal/app"
	"github.com/vnchk1/inventory-control/internal/config"
	logging "github.com/vnchk1/inventory-control/internal/logger"
)

const (
	ShutdownTimeoutSeconds = 5
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
		log.Fatalf("Error loading config: %v\n", err)
	}

	logger := logging.NewLogger(cfg.LogLevel)

	app, err := inventoryApp.NewApp(cfg, logger)
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

	err = app.Stop(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Warn("Shutdown timed out - some resources may not be fully released")
		} else {
			logger.Error("Shutdown failed:", err)
		}
	} else {
		logger.Info("Graceful shutdown completed")
	}
}
