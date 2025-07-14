package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

type Products struct {
	Id         int
	Name       string
	Price      int
	Quantity   int
	CategoryId int
}

func CreateProduct(conn *pgx.Conn, product *Products) error {
	query := `
	INSERT INTO products (product_name, price, quantity, category_id) 
	VALUES ($1, $2, $3, $4) RETURNING product_id`

	return conn.QueryRow(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Quantity,
		product.CategoryId).Scan(&product.Id)
}

func InitDB() (*pgx.Conn, error) {
	ConnStr := "user=postgres password=postgres host=localhost port=5432 dbname=inventory_control sslmode=disable"
	conn, err := pgx.Connect(context.Background(), ConnStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return conn, nil
}

func main() {
	conn, err := InitDB()
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
