package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/vnchk1/inventory-control/internal/config"
	"testing"
	"time"
)

var (
	migPath = "C:/Users/user/GolandProjects/inventory-control/migrations"
)

func SetupTestContainer(t *testing.T) (*DB, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "inventory_control",
		},
		WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
			return fmt.Sprintf(
				"host=%s port=%s user=postgres password=postgres dbname=inventory_control sslmode=disable",
				host, port.Port(),
			)
		}).WithStartupTimeout(60 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	host, err := container.Host(ctx)
	require.NoError(t, err)
	port, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	cfg := &config.Config{
		Log:    nil,
		Server: nil,
		DB: &config.DatabaseConfig{
			Host:     host,
			Port:     port.Port(),
			Username: "postgres",
			Password: "postgres",
			DBName:   "inventory_control",
			SSLMode:  "disable",
		},
		Migrator: nil,
	}

	testDB, err := NewDB(cfg)
	require.NoError(t, err)

	sqlDB, err := sql.Open("postgres", testDB.GetConnString(cfg))
	require.NoError(t, err)
	defer sqlDB.Close()

	err = goose.SetDialect("postgres")
	require.NoError(t, err)
	err = goose.Up(sqlDB, migPath)
	require.NoError(t, err)

	cleanup := func() {
		testDB.Close()
		_ = container.Terminate(ctx)
	}

	return testDB, cleanup
}
