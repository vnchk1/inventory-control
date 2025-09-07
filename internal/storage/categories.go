package storage

import (
	"context"
	"fmt"

	"github.com/vnchk1/inventory-control/internal/models"
)

type CategoryStorage struct {
	Pool *DB
}

func (c *CategoryStorage) Read(ctx context.Context, id int) (category models.Category, err error) {
	query := `
	SELECT * FROM categories WHERE category_id = $1`

	err = c.Pool.QueryRow(
		ctx, query,
		id).Scan(&category.ID,
		&category.Name)
	if err != nil {
		err = fmt.Errorf("row SELECT error: %w", err)

		return
	}

	return
}

func (c *CategoryStorage) Update(ctx context.Context, category models.Category) (err error) {
	query := `UPDATE categories SET 
		category_name = $1
		WHERE category_id = $2`

	_, err = c.Pool.Exec(ctx, query,
		category.Name, category.ID)
	if err != nil {
		return fmt.Errorf("row UPDATE error: %w", err)
	}

	return
}

func (c *CategoryStorage) Delete(ctx context.Context, id int) (err error) {
	query := `DELETE FROM categories WHERE category_id = $1`

	result, err := c.Pool.Exec(ctx, query, id)
	if err != nil {
		err = fmt.Errorf("row DELETE error: %w", err)

		return
	}

	if result.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return
}

func NewCategoryStorage(pool *DB) *CategoryStorage {
	return &CategoryStorage{Pool: pool}
}

func (c *CategoryStorage) Create(ctx context.Context, category *models.Category) (err error) {
	query := `
	INSERT INTO categories (category_name)
	VALUES ($1) RETURNING category_id`

	err = c.Pool.QueryRow(
		ctx, query,
		category.Name,
	).Scan(&category.ID)

	if err != nil {
		return fmt.Errorf("row INSERT error: %w", err)
	}

	return
}
