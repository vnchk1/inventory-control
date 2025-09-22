package server

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/vnchk1/inventory-control/internal/models"
)

//go:generate mockgen -source=categories.go -destination=../mocks/category_usecase_mock.go -package=mocks
type (
	CategoryUseCase interface {
		Create(ctx context.Context, category *models.Category) error
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

//nolint:varnamelen
func (p *CategoryHandler) Create(c echo.Context) error {
	var req models.Category

	err := c.Bind(&req)
	if err != nil {
		p.Logger.Error("error parsing JSON", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err = p.Service.Create(c.Request().Context(), &req)

	if err != nil {
		p.Logger.Error("services.categories.Create", "error", err)

		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}

//nolint:varnamelen
func (p *CategoryHandler) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.Logger.Error("Invalid ID", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	product, err := p.Service.Read(c.Request().Context(), id)
	if err != nil {
		p.Logger.Error("services.categories.Read", "error", err)

		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, product)
}

//nolint:varnamelen
func (p *CategoryHandler) Update(c echo.Context) error {
	var req models.Category

	err := c.Bind(&req)
	if err != nil {
		p.Logger.Error("error parsing JSON", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err = p.Service.Update(c.Request().Context(), req)
	if err != nil {
		p.Logger.Error("services.categories.Update", "error", err)

		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}

//nolint:varnamelen
func (p *CategoryHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.Logger.Error("Invalid ID", "error", err)

		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err = p.Service.Delete(c.Request().Context(), id)
	if err != nil {
		p.Logger.Error("services.categories.Delete", "error", err)

		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	return c.JSON(http.StatusNoContent, http.StatusText(http.StatusNoContent))
}
