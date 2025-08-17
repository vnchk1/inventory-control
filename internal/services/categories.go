package services

import (
	"context"
	"fmt"

	"github.com/vnchk1/inventory-control/internal/models"
)

type CategoryRepo interface {
	Create(ctx context.Context, category models.Category) error
	Read(ctx context.Context, id int) (models.Category, error)
	Update(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryService struct {
	Storage CategoryRepo
}

func (c *CategoryService) Read(ctx context.Context, id int) (category models.Category, err error) {
	category, err = c.Storage.Read(ctx, id)
	if err != nil {
		return models.Category{}, fmt.Errorf("storage.categories.Read: %w", err)
	}

	return
}

func (c *CategoryService) Update(ctx context.Context, category models.Category) error {
	// TODO implement me
	panic("implement me")
}

func (c *CategoryService) Delete(ctx context.Context, id int) error {
	// TODO implement me
	panic("implement me")
}

func NewCategoryService(storage CategoryRepo) *CategoryService {
	return &CategoryService{
		Storage: storage,
	}
}

func (c *CategoryService) Create(ctx context.Context, category models.Category) (err error) {
	if category.Name == "" {
		return models.NewEmptyErr("name")
	}

	err = c.Storage.Create(ctx, category)

	if err != nil {
		return fmt.Errorf("storage.categories.Create: %w", err)
	}

	return
}
