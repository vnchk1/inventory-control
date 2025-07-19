package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/vnchk1/inventory_control/config"
	"log"
	"github.com/vnchk1/inventory_control/internal/models"
)

func ConnStr(cfg *config.Config) string {
	return fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}

func InitDB(ConnStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), ConnStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return conn, nil
}

func CreateProduct(conn *pgx.Conn, product *Products) error {
	query := `
	INSERT INTO products (product_name, price, quantity, category_id) 
	VALUES ($1, $2, $3, $4) RETURNING product_id`

	return conn.QueryRow(
		context.Background(), +query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryId).Scan(&product.Id)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	conn, err := InitDB(ConnStr(cfg))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	user := &Products{
		Id:         1,
		Name:       "guest",
		Price:      1000,
		Quantity:   1,
		CategoryId: 1,
	}

	err = CreateProduct(conn, user)
	if err != nil {
		log.Fatalf("Unable to create product: %v\n", err)
	}
}
