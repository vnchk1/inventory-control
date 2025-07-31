package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/inventory-control/internal/models"
	"log/slog"
	"net/http"
)

type (
	CategoryUseCase interface {
		Create(ctx context.Context, category models.Category) error
		Read(ctx context.Context, id int) (models.Category, error)
		Update(ctx context.Context, category models.Category) error
		Delete(ctx context.Context, id int) error
	}

	CategoryHandler struct {
		Service CategoryUseCase
		Logger  *slog.Logger
	}
)

func NewCategoryHandler(uc CategoryUseCase, logger *slog.Logger) *CategoryHandler {
	return &CategoryHandler{
		Service: uc,
		Logger:  logger,
	}
}

func (p *CategoryHandler) Create(c echo.Context) error {
	var req models.Category

	err := c.Bind(&req)
	if err != nil {
		p.Logger.Error("error parsing JSON", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err = p.Service.Create(c.Request().Context(), req)
	if err != nil {
		p.Logger.Error("services.categories.Create", "error", err)

		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, req)
}
