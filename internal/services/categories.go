package services

import (
	"context"
	"fmt"

	"github.com/vnchk1/inventory-control/internal/models"
)

type CategoryRepo interface {
	Create(ctx context.Context, category *models.Category) error
	Read(ctx context.Context, id int) (models.Category, error)
	Update(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryService struct {
	Storage CategoryRepo
}

func NewCategoryService(storage CategoryRepo) *CategoryService {
	return &CategoryService{
		Storage: storage,
	}
}

func (c *CategoryService) Read(ctx context.Context, id int) (category models.Category, err error) {
	category, err = c.Storage.Read(ctx, id)
	if err != nil {
		return models.Category{}, fmt.Errorf("storage.categories.Read: %w", err)
	}

	return
}

func (c *CategoryService) Update(ctx context.Context, category models.Category) (err error) {
	if category.ID <= 0 {
		return models.NewEmptyErr("id")
	}

	if category.Name == "" {
		return models.NewEmptyErr("name")
	}

	err = c.Storage.Update(ctx, category)
	if err != nil {
		return fmt.Errorf("storage.categories.Update: %w", err)
	}

	return
}

func (c *CategoryService) Delete(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.NewNegativeErr("id")
	}

	err = c.Storage.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("storage.categories.Delete: %w", err)
	}

	return
}

func (c *CategoryService) Create(ctx context.Context, category *models.Category) (err error) {
	if category.Name == "" {
		return models.NewEmptyErr("name")
	}

	err = c.Storage.Create(ctx, category)

	if err != nil {
		return fmt.Errorf("storage.categories.Create: %w", err)
	}

	return
}
