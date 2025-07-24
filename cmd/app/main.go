package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/inventory-control/internal/config"
	logging "github.com/vnchk1/inventory-control/internal/logger"
	"github.com/vnchk1/inventory-control/internal/middleware"
	"github.com/vnchk1/inventory-control/internal/server"
	"github.com/vnchk1/inventory-control/internal/services/products"
	"github.com/vnchk1/inventory-control/internal/storage"
	"github.com/vnchk1/inventory-control/internal/storage/db"
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
	//логгер
	logger := logging.NewLogger(cfg.LogLevel)
	//БД
	pool, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v\n", err)
	}
	//работа с БД
	productStorage, err := storage.NewProducts(pool)
	if err != nil {
		log.Fatalf("Error creating product storage: %v\n", err)
	}
	//сервис
	productService := products.NewProductService(productStorage)
	//продуктовый хендлер
	//productHandler := server.NewProductHandler(productService, logger)
	//хендлеры
	h := server.NewHandlers(productService, logger)
	//err = server.NewServer(cfg)

	e := echo.New()
	e.Use(middleware.LoggingMiddleware(logger))
	productGroup := e.Group("/products")
	productGroup.GET("/:id", h.Products.Read)

	err = e.Start(":" + cfg.ServerPort)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

}
