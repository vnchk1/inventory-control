package storage

import (
	"context"
	"fmt"
	"github.com/vnchk1/inventory-control/internal/models"
)

type ProductStorage struct {
	pool *DB
}

func NewProductStorage(db *DB) *ProductStorage {
	return &ProductStorage{pool: db}
}

func (p *ProductStorage) Create(ctx context.Context, product models.Product) (err error) {
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
		return fmt.Errorf("row INSERT error: %w", err)
	}

	return
}

func (p *ProductStorage) Read(ctx context.Context, id int) (product models.Product, err error) {
	query := `
	SELECT * FROM products WHERE product_id = $1`

	err = p.pool.QueryRow(
		ctx, query,
		id).Scan(&product.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.CategoryID)
	if err != nil {
		err = fmt.Errorf("row SELECT error: %w", err)
		return
	}

	return
}

func (p *ProductStorage) Update(ctx context.Context, product models.Product) (err error) {
	query := `UPDATE products SET 
		product_name = $1, 
		price = $2, 
		quantity = $3, 
		category_id = $4
		WHERE product_id = $5`

	_, err = p.pool.Exec(ctx, query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryID,
		product.ID)
	if err != nil {
		return fmt.Errorf("row UPDATE error: %w", err)
	}

	return
}

func (p *ProductStorage) Delete(ctx context.Context, id int) (err error) {
	query := `DELETE FROM products WHERE product_id = $1`

	result, err := p.pool.Exec(ctx, query, id)
	if err != nil {
		err = fmt.Errorf("row DELETE error: %w", err)
		return
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return
}
