package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/inventory-control/internal/models"
	"log/slog"
	"net/http"
	"strconv"
)

type (
	ProductService interface {
		Create(context.Context, models.Product) error
		Read(context.Context, int) (models.Product, error)
		Update(context.Context, models.Product) error
		Delete(context.Context, int) error
	}

	ProductHandler struct {
		Service ProductService
		Logger  *slog.Logger
	}
)

func NewProductHandler(service ProductService, logger *slog.Logger) *ProductHandler {
	return &ProductHandler{
		Service: service,
		Logger:  logger,
	}
}

func (p *ProductHandler) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.Logger.Error("Invalid ID", "error", err)
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	product, err := p.Service.Read(c.Request().Context(), id)
	if err != nil {
		p.Logger.Error("services.products.Read", "error", err)
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}
