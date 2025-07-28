APP_NAME := inventory-control
MAIN_PATH := cmd/inventory-control/main.go
MIGRATOR_PATH := cmd/migrator/main.go

.PHONY: all build run swag fmt lint test clean tidy

all: build

build:
	@echo "Running build..."
	go build -o $(APP_NAME) $(MAIN_PATH)

run:
	@echo "Running the app..."
	go run $(MAIN_PATH)

migrate:
	@echo "Running migrations"
	go run $(MIGRATOR_PATH)

goosedown:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/inventory_control?sslmode=disable down


lint:
	golangci-lint run


swag:
	swag init -g $(MAIN_PATH) --output docs

fmt:
	go fmt ./...

test:
	go test ./... -v

clean:
	go clean
	if exist $(APP_NAME) del $(APP_NAME)

tidy:
	go mod tidy