package server

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vnchk1/inventory-control/internal/models"
)

type (
	ProductService interface {
		Create(ctx context.Context, product models.Product) error
		Read(ctx context.Context, id int) (models.Product, error)
		Update(ctx context.Context, product models.Product) error
		Delete(ctx context.Context, id int) error
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

func (p *ProductHandler) Create(c echo.Context) error {
	var req models.Product

	err := c.Bind(&req)
	if err != nil {
		p.Logger.Error("error parsing JSON", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err = p.Service.Create(c.Request().Context(), req)
	if err != nil {
		p.Logger.Error("services.product.Create", "error", err)

		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, req)
}

func (p *ProductHandler) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.Logger.Error("Invalid ID", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	product, err := p.Service.Read(c.Request().Context(), id)
	if err != nil {
		p.Logger.Error("services.products.Read", "error", err)

		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) Update(c echo.Context) error {
	var req models.Product

	err := c.Bind(&req)
	if err != nil {
		p.Logger.Error("error parsing JSON", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err = p.Service.Update(c.Request().Context(), req)
	if err != nil {
		p.Logger.Error("services.product.Update", "error", err)

		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, req)
}

func (p *ProductHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.Logger.Error("Invalid ID", "error", err)
	}

	err = p.Service.Delete(c.Request().Context(), id)
	if err != nil {
		p.Logger.Error("services.products.Delete", "error", err)

		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	return c.JSON(http.StatusNoContent, nil)
}
