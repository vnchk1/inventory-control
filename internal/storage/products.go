package storage

import (
	"context"
	"fmt"
	"github.com/vnchk1/inventory-control/internal/db"
	"github.com/vnchk1/inventory-control/internal/models"
)

type Products struct {
	pool *db.DB
}

func NewProducts(db *db.DB) (*Products, error) {
	return &Products{pool: db}, nil
}

func (p *Products) Create(ctx context.Context, product *models.Product) (err error) {
	query := `
	INSERT INTO products (product_name, price, quantity, category_id) 
	VALUES ($1, $2, $3, $4) RETURNING product_id`

	err = p.pool.QueryRow(
		ctx, query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryID).Scan(&product.ID)
	if err != nil {
		err = fmt.Errorf("error inserting row: %w", err)
		return
	}

	//logger.Debug("Product created",
	//	"product_name", product.Name,
	//	"price", product.Price,
	//	"quantity", product.Quantity,
	//	"category_id", product.CategoryID,
	//	"product_id", product.ID)

	return
}

func (p *Products) Read(ctx context.Context, id int) (err error) {
	var product models.Product

	query := `
	SELECT * FROM products WHERE product_id = $1`

	err = p.pool.QueryRow(
		context.Background(), query,
		id).Scan(&product.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.CategoryID)
	if err != nil {
		err = fmt.Errorf("error selecting row: %w", err)
		return
	}

	//logger.Debug("Product read",
	//	"product_id", product.ID,
	//	"product_name", product.Name,
	//	"price", product.Price,
	//	"quantity", product.Quantity,
	//	"category_id", product.CategoryID)

	return
}

func (p *Products) Update(ctx context.Context, product *models.Product) (err error) {
	query := `UPDATE products SET 
		product_name = $1, 
		price = $2, 
		quantity = $3, 
		category_id = $4
		WHERE product_id = $5`

	_, err = p.pool.Exec(context.Background(), query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryID,
		product.ID)
	if err != nil {
		err = fmt.Errorf("error updating row: %w", err)
		return
	}

	//logger.Debug("Product updated",
	//	"product_id", product.ID,
	//	"product_name", product.Name,
	//	"price", product.Price,
	//	"quantity", product.Quantity,
	//	"category_id", product.CategoryID)

	return
}

func (p *Products) Delete(ctx context.Context, id int) (err error) {
	query := `DELETE FROM products WHERE product_id = $1`

	result, err := p.pool.Exec(context.Background(), query, id)
	if err != nil {
		err = fmt.Errorf("error deleting row: %w", err)
		return
	}

	if result.RowsAffected() == 0 {
		err = fmt.Errorf("row not found")
		return
	}

	//logger.Debug("Product deleted",
	//	"product_id", id)

	return
}
