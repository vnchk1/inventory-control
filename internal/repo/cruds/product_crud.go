package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/vnchk1/inventory-control/internal/models"
	"log/slog"
)

func Create(conn *pgx.Conn, logger *slog.Logger, product *models.Products) (err error) {
	query := `
	INSERT INTO products (product_name, price, quantity, category_id) 
	VALUES ($1, $2, $3, $4) RETURNING product_id`

	err = conn.QueryRow(
		context.Background(), query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryId).Scan(&product.Id)

	logger.Info("Product created",
		"product_name", product.Name,
		"price", product.Price,
		"quantity", product.Quantity,
		"category_id", product.CategoryId,
		"product_id", product.Id)

	return
}

func Read(conn *pgx.Conn, logger *slog.Logger, id int) (err error) {
	var product models.Products

	query := `
	SELECT * FROM products WHERE product_id = $1`

	err = conn.QueryRow(
		context.Background(), query,
		id).Scan(&product.Id,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.CategoryId)

	logger.Info("Product read",
		"product_id", product.Id,
		"product_name", product.Name,
		"price", product.Price,
		"quantity", product.Quantity,
		"category_id", product.CategoryId)

	return
}

func Update(conn *pgx.Conn, logger *slog.Logger, product *models.Products) (err error) {
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
		product.CategoryId,
		product.Id)

	logger.Info("Product updated",
		"product_id", product.Id,
		"product_name", product.Name,
		"price", product.Price,
		"quantity", product.Quantity,
		"category_id", product.CategoryId)

	return
}
