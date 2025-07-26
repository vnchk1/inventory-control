package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/inventory-control/internal/config"
	"github.com/vnchk1/inventory-control/internal/middleware"
	"log/slog"
	"net/http"
)

type Server struct {
	Echo   *echo.Echo
	Logger *slog.Logger
	Config *config.Config
}

func NewServer(cfg *config.Config, logger *slog.Logger) *Server {
	e := echo.New()

	e.Use(middleware.LoggingMiddleware(logger))

	return &Server{
		Echo:   e,
		Logger: logger,
		Config: cfg,
	}
}

func (s *Server) RegisterRoutes(h *Handlers) {
	productGroup := s.Echo.Group("/products")

	productGroup.GET("/:id", h.Products.Read)
	productGroup.POST("/create", h.Products.Create)
	productGroup.PUT("/update", h.Products.Update)
	productGroup.DELETE("/:id", h.Products.Delete)
}

func (s *Server) Start() (err error) {
	go func() {
		err = s.Echo.Start(":" + s.Config.ServerPort)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error("server.Start: ", "error", err, "port", s.Config.ServerPort)
		}
	}()
	return
}
