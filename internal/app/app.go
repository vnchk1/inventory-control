package app

import (
	"context"
	"github.com/vnchk1/inventory-control/internal/config"
	"github.com/vnchk1/inventory-control/internal/server"
	productservice "github.com/vnchk1/inventory-control/internal/services/products"
	"github.com/vnchk1/inventory-control/internal/storage"
	"log/slog"
)

type App struct {
	Server *server.Server
	DB     *storage.DB
	Logger *slog.Logger
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	//инициализация БД
	pool, err := storage.NewDB(cfg)
	if err != nil {
		logger.Error("Error connecting to DB: %v\n", "error", err)
	}
	//работа с БД
	productStorage := storage.NewProductStorage(pool)
	//use cases
	productService := productservice.NewProductService(productStorage)
	//infrastructure
	handlers := server.NewHandlers(productService, logger)
	//инициализация сервера
	newServer := server.NewServer(cfg, logger)
	//регистрация маршрутов
	newServer.RegisterRoutes(handlers)
	return &App{
		Server: newServer,
		DB:     pool,
		Logger: logger,
	}
}

func (p *App) Run() {
	err := p.Server.Run()
	if err != nil {
		p.Logger.Error("Error starting server: %v\n", "error", err)
	}
}

func (p *App) Stop(ctx context.Context) {
	p.Logger.Info("DB pool closing...")
	p.DB.Close(ctx)

	doneChan := make(chan error)
	go func() {
		doneChan <- p.Server.Stop(ctx)
	}()

	select {
	case err := <-doneChan:
		if err != nil {
			p.Logger.Error("server.Stop: ", "error", err)
		}
		p.Logger.Info("App shut down...")
	case <-ctx.Done():
		p.Logger.Warn("App stoped forced by timeout")
	}
}
