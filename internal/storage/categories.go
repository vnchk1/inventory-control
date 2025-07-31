package storage

import (
	"context"
	"fmt"

	"github.com/vnchk1/inventory-control/internal/models"
)

type CategoryStorage struct {
	pool *DB
}

func (c *CategoryStorage) Read(ctx context.Context, id int) (models.Category, error) {
	// TODO implement me
	panic("implement me")
}

func (c *CategoryStorage) Update(ctx context.Context, category models.Category) error {
	// TODO implement me
	panic("implement me")
}

func (c *CategoryStorage) Delete(ctx context.Context, id int) error {
	// TODO implement me
	panic("implement me")
}

func NewCategoryStorage(pool *DB) *CategoryStorage {
	return &CategoryStorage{pool: pool}
}

func (c *CategoryStorage) Create(ctx context.Context, category models.Category) (err error) {
	query := `
	INSERT INTO categories (category_name)
	VALUES ($1) RETURNING category_id`

	err = c.pool.QueryRow(
		ctx, query,
		category.Name,
	).Scan(&category.ID)
	if err != nil {
		return fmt.Errorf("row INSERT error: %w", err)
	}

	return
}
