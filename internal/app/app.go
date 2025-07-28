package app

import (
	"context"
	"github.com/vnchk1/inventory-control/internal/config"
	server1 "github.com/vnchk1/inventory-control/internal/server"
	productservice "github.com/vnchk1/inventory-control/internal/services/products"
	"github.com/vnchk1/inventory-control/internal/storage"
	"log/slog"
)

type App struct {
	Server *server1.Server
	DB     *storage.DB
	Logger *slog.Logger
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	//инициализация БД
	pool, err := storage.NewDB(cfg)
	if err != nil {
		logger.Error("Error connecting to DB: %v\n", "error", err)
	}
	logger.Info("Connected to DB", "stat", pool.GetConnString())
	//работа с БД
	productStorage := storage.NewProductStorage(pool)
	//use cases
	productService := productservice.NewProductService(productStorage)
	//infrastructure
	handlers := server1.NewHandlers(productService, logger)
	//инициализация сервера
	server := server1.NewServer(cfg, logger)
	logger.Info("Starting server", "port", server.Config.ServerPort)
	//регистрация маршрутов
	server.RegisterRoutes(handlers)
	return &App{
		Server: server,
		DB:     pool,
		Logger: logger,
	}
}

func (p *App) Run() (err error) {
	err = p.Server.Run()
	if err != nil {
		p.Logger.Error("app.Run: %v\n", "error", err)
		return
	}
	return
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
