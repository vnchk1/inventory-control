package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vnchk1/inventory-control/internal/config"
	logging "github.com/vnchk1/inventory-control/internal/logger"
	srv "github.com/vnchk1/inventory-control/internal/server"
	productservice "github.com/vnchk1/inventory-control/internal/services/products"
	"github.com/vnchk1/inventory-control/internal/storage"
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
	//инициализация БД
	pool, err := storage.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v\n", err)
	}
	//работа с БД
	productStorage := storage.NewProductStorage(pool)
	//use cases
	productService := productservice.NewProductService(productStorage)
	//infrastructure
	handlers := srv.NewHandlers(productService, logger)
	//инициализация сервера
	server := srv.NewServer(cfg, logger)
	//регистрация маршрутов
	server.RegisterRoutes(handlers)
	//старт сервера
	err = server.Start()
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//освобождение ресурсов
	logger.Info("DB pool closing...")
	pool.Close(ctx)

	err = server.Echo.Shutdown(ctx)
	if err != nil {
		logger.Error("server forced to shutdown", "error", err)
	} else {
		logger.Info("server gracefully shutdown")
	}
}
