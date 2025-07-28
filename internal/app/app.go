package app

import (
	"context"
	"log/slog"

	"github.com/vnchk1/inventory-control/internal/config"
	serverpkg "github.com/vnchk1/inventory-control/internal/server"
	productservice "github.com/vnchk1/inventory-control/internal/services"
	"github.com/vnchk1/inventory-control/internal/storage"
)

type App struct {
	Server *serverpkg.Server
	DB     *storage.DB
	Logger *slog.Logger
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	pool, err := storage.NewDB(cfg)
	if err != nil {
		logger.Error("Error connecting to DB: %v\n", "error", err)
	}

	logger.Info("Connected to DB", "conn string", pool.GetConnString())

	productStorage := storage.NewProductStorage(pool)
	productService := productservice.NewProductService(productStorage)
	handlers := serverpkg.NewHandlers(productService, logger)

	server := serverpkg.NewServer(cfg, logger)
	logger.Info("Starting server", "port", server.Config.ServerPort)

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
		doneChan <- p.Server.Stop()
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
