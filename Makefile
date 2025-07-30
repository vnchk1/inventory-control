APP_NAME := inventory-control
MIGRATOR_NAME := migrator
MAIN_PATH := ./cmd/inventory-control
MIGRATOR_PATH := ./cmd/migrator

.PHONY: all build run swag fmt lint lint-fix test tidy docker-build docker-testrun docker-run mdocker-build mdocker-run migrate-build migrate-run

all: build

build:
	@echo "Running build..."
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

run:
	@echo "Running the app..."
	go run $(MAIN_PATH)

migrate-build:
	@echo "Building migrations..."
	go build -o bin/$(MIGRATOR_NAME) $(MIGRATOR_PATH)

migrate-run:
	@echo "Running migrations"
	go run $(MIGRATOR_PATH)

goosedown:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/inventory_control?sslmode=disable down

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

swag:
	swag init -g $(MAIN_PATH) --output docs

fmt:
	go fmt ./...

test:
	go test ./... -v

tidy:
	go mod tidy

docker-build:
	docker build -t inventory-control:latest -f Dockerfile.inventory-control .

docker-testrun:
	docker run --rm -it inventory-control:latest

docker-run:
	docker run --rm -p 8080:8080 -e "CONFIG_PATH=configs/.env" --name inventory-control inventory-control:latest

mdocker-build:
	docker build -t migrator:latest -f Dockerfile.migrator .

mdocker-run:
	docker run --rm -p 8080:8080 -e "CONFIG_PATH=configs/.env" --name migrator migrator:latest
