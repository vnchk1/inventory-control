package services

import (
	"context"
	"fmt"

	"github.com/vnchk1/inventory-control/internal/models"
)

type ProductRepo interface {
	Create(ctx context.Context, product models.Product) error
	Read(ctx context.Context, id int) (models.Product, error)
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, id int) error
}

type Products struct {
	Storage ProductRepo
}

func NewProductService(storage ProductRepo) *Products {
	return &Products{
		Storage: storage,
	}
}

func (c *Products) Create(ctx context.Context, product models.Product) (err error) {
	if product.Name == "" {
		return models.NewEmptyErr("name")
	}

	if product.Price < 0 {
		return models.NewNegativeErr("price")
	}

	if product.Quantity < 0 {
		return models.NewNegativeErr("quantity")
	}

	err = c.Storage.Create(ctx, product)
	if err != nil {
		return fmt.Errorf("storage.products.Create: %w", err)
	}

	return
}

func (c *Products) Update(ctx context.Context, product models.Product) (err error) {
	if product.ID <= 0 {
		return models.NewEmptyErr("id")
	}

	if product.Name == "" {
		return models.NewEmptyErr("name")
	}

	if product.Price < 0 {
		return models.NewNegativeErr("price")
	}

	if product.Quantity < 0 {
		return models.NewNegativeErr("quantity")
	}

	if product.CategoryID <= 0 {
		return models.NewEmptyErr("category_id")
	}

	err = c.Storage.Update(ctx, product)
	if err != nil {
		return fmt.Errorf("stoeage.products.Update: %w", err)
	}

	return
}

func (c *Products) Delete(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.NewNegativeErr("id")
	}

	err = c.Storage.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("storage.products.Delete: %w", err)
	}

	return
}

func (c *Products) Read(ctx context.Context, id int) (product models.Product, err error) {
	product, err = c.Storage.Read(ctx, id)
	if err != nil {
		return models.Product{}, fmt.Errorf("storage.products.Read: %w", err)
	}

	return
}
