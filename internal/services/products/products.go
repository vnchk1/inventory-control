package products

import (
	"context"
	"fmt"
	"github.com/vnchk1/inventory-control/internal/models"
)

type ProductRepo interface {
	Create(context.Context, models.Product) error
	Read(context.Context, int) (models.Product, error)
	Update(context.Context, models.Product) error
	Delete(context.Context, int) error
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
		return fmt.Errorf("name cannot be empty %v", product.Name)
	}
	if product.Price < 0 {
		return fmt.Errorf("price cannot be negative %v", product.Price)
	}
	if product.Quantity < 0 {
		return fmt.Errorf("quantity cannot be negative %v", product.Quantity)
	}
	if product.CategoryID <= 0 {
		return fmt.Errorf("category_id cannot be negative %v", product.CategoryID)
	}

	err = c.Storage.Create(ctx, product)
	if err != nil {
		return fmt.Errorf("storage.products.Create: %w", err)
	}
	return
}

func (c *Products) Update(ctx context.Context, product models.Product) (err error) {
	if product.ID <= 0 {
		return fmt.Errorf("product_id must be a positive number %v", product.ID)
	}
	if product.Name == "" {
		return fmt.Errorf("name cannot be empty %v", product.Name)
	}
	if product.Price < 0 {
		return fmt.Errorf("price cannot be negative %v", product.Price)
	}
	if product.Quantity < 0 {
		return fmt.Errorf("quantity cannot be negative %v", product.Quantity)
	}
	if product.CategoryID <= 0 {
		return fmt.Errorf("category_id cannot be negative %v", product.CategoryID)
	}

	err = c.Storage.Update(ctx, product)
	if err != nil {
		return fmt.Errorf("stoeage.products.Update: %w", err)
	}
	return
}

func (c *Products) Delete(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return fmt.Errorf("product_id must be a positive number %v", id)
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
