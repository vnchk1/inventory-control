package app

import (
	"context"
	"errors"
	"log/slog"

	"github.com/vnchk1/inventory-control/internal/config"
	"github.com/vnchk1/inventory-control/internal/server"
	"github.com/vnchk1/inventory-control/internal/services"
	"github.com/vnchk1/inventory-control/internal/storage"
)

var ErrDBConnectionFailed = errors.New("DB connection failed")

type App struct {
	Server *server.Server
	DB     *storage.DB
	Logger *slog.Logger
}

func NewApp(cfg *config.Config, logger *slog.Logger) (*App, error) {
	pool, err := storage.NewDB(cfg)
	if err != nil {
		logger.Error("Error connecting to DB: %v\n", "error", err)

		return nil, ErrDBConnectionFailed
	}

	logger.Debug("Connected to DB", "conn string", pool.GetConnString(cfg))

	categoryStorage := storage.NewCategoryStorage(pool)
	categoryService := services.NewCategoryService(categoryStorage)

	productStorage := storage.NewProductStorage(pool)
	productService := services.NewProductService(productStorage)

	handlers := server.NewHandlers(categoryService, productService, logger)

	newServer := server.NewServer(cfg, logger)
	logger.Debug("Starting server", "port", newServer.Config.ServerPort)

	newServer.RegisterRoutes(handlers)

	return &App{
		Server: newServer,
		DB:     pool,
		Logger: logger,
	}, nil
}

func (p *App) Run() (err error) {
	err = p.Server.Run()
	if err != nil {
		p.Logger.Error("app.Run: %v\n", "error", err)

		return
	}

	return
}

func (p *App) Stop(ctx context.Context) (err error) {
	p.Logger.Info("Server stopping...")
	if err = p.Server.Stop(ctx); err != nil {
		p.Logger.Error("app.Stop: ", "error", err)
		return
	}

	p.Logger.Info("DB pool closing...")
	dbClosed := make(chan struct{})
	go func() {
		defer close(dbClosed)
		p.DB.Close()
	}()

	select {
	case <-dbClosed:
		p.Logger.Info("DB closed successfully")
	case <-ctx.Done():
		p.Logger.Warn("DB close interrupted by context timeout")
		return ctx.Err()
	}

	return nil
}
