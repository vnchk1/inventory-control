package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/vnchk1/inventory-control/internal/models"
	"log/slog"
)

func Create(conn *pgx.Conn, product *models.Products, logger *slog.Logger) error {
	query := `
	INSERT INTO products (product_name, price, quantity, category_id) 
	VALUES ($1, $2, $3, $4) RETURNING product_id`

	logger.Info("Product created",
		"Name", product.Name,
		"Price", product.Price,
		"Quantity", product.Quantity,
		"Category_ID", product.CategoryId)
	return conn.QueryRow(
		context.Background(), query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryId).Scan(&product.Id)
}
