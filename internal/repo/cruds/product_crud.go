package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/vnchk1/inventory-control/internal/models"
	"log/slog"
)

func Create(conn *pgx.Conn, logger *slog.Logger, product *models.Product) (err error) {
	query := `
	INSERT INTO products (product_name, price, quantity, category_id) 
	VALUES ($1, $2, $3, $4) RETURNING product_id`

	err = conn.QueryRow(
		context.Background(), query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryID).Scan(&product.ID)

	logger.Debug("Product created",
		"product_name", product.Name,
		"price", product.Price,
		"quantity", product.Quantity,
		"category_id", product.CategoryID,
		"product_id", product.ID)

	return
}

func Read(conn *pgx.Conn, logger *slog.Logger, id int) (err error) {
	var product models.Product

	query := `
	SELECT * FROM products WHERE product_id = $1`

	err = conn.QueryRow(
		context.Background(), query,
		id).Scan(&product.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.CategoryID)

	logger.Debug("Product read",
		"product_id", product.ID,
		"product_name", product.Name,
		"price", product.Price,
		"quantity", product.Quantity,
		"category_id", product.CategoryID)

	return
}

func Update(conn *pgx.Conn, logger *slog.Logger, product *models.Product) (err error) {
	query := `UPDATE products SET 
		product_name = $1, 
		price = $2, 
		quantity = $3, 
		category_id = $4
		WHERE product_id = $5`

	_, err = conn.Exec(context.Background(), query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryID,
		product.ID)

	logger.Debug("Product updated",
		"product_id", product.ID,
		"product_name", product.Name,
		"price", product.Price,
		"quantity", product.Quantity,
		"category_id", product.CategoryID)

	return
}
