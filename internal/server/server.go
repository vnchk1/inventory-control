package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnchk1/inventory-control/internal/config"
	"github.com/vnchk1/inventory-control/internal/middleware"
)

type Server struct {
	Echo   *echo.Echo
	Logger *slog.Logger
	Config *config.Config
}

func NewServer(cfg *config.Config, logger *slog.Logger) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.LoggingMiddleware(logger))

	return &Server{
		Echo:   e,
		Logger: logger,
		Config: cfg,
	}
}

func (s *Server) RegisterRoutes(h *Handlers) {
	categoryGroup := s.Echo.Group("/categories")

	categoryGroup.GET("/:id", h.Categories.Read)
	categoryGroup.POST("/create", h.Categories.Create)
	categoryGroup.PUT("/update", h.Categories.Update)
	categoryGroup.DELETE("/:id", h.Categories.Delete)

	productGroup := s.Echo.Group("/products")

	productGroup.GET("/:id", h.Products.Read)
	productGroup.POST("/create", h.Products.Create)
	productGroup.PUT("/update", h.Products.Update)
	productGroup.DELETE("/:id", h.Products.Delete)
}

func (s *Server) Run() (err error) {
	go func() {
		err = s.Echo.Start(":" + s.Config.ServerPort)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error("server.Run: ", "error", err, "port", s.Config.ServerPort)
		}
	}()

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	err = s.Echo.Shutdown(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.Logger.Error("server.Stop: ", "error", err, "port", s.Config.ServerPort)
	}
	s.Logger.Info("Server stopped")

	return
}
