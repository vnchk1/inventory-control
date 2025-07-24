package products

import (
	"context"
	"fmt"
	"github.com/vnchk1/inventory-control/internal/models"
)

type (
	ProductRepo interface {
		Create(context.Context, models.Product) error
		Read(context.Context, int) (models.Product, error)
		Update(context.Context, models.Product) error
		Delete(context.Context, int) error
	}

	Products struct {
		Storage ProductRepo
	}
)

func NewProductService(storage ProductRepo) *Products {
	return &Products{
		Storage: storage,
	}
}

func (c *Products) Create(ctx context.Context, product models.Product) error {
	//TODO implement me
	panic("implement me")
}

func (c *Products) Update(ctx context.Context, product models.Product) error {
	//TODO implement me
	panic("implement me")
}

func (c *Products) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (c *Products) Read(ctx context.Context, id int) (product models.Product, err error) {
	product, err = c.Storage.Read(ctx, id)
	if err != nil {
		err = fmt.Errorf("Products.Storage.Read: %w", err)
		return
	}
	return
}
