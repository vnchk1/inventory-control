package config

import (
	"fmt"
	"os"

	"github.com/vnchk1/inventory-control/internal/models"
)

type MigratorConfig struct {
	MigrationsPath string
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	DBSSLMode      string
}

func LoadMigratorConfig() (*MigratorConfig, error) {
	config := &MigratorConfig{}

	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.DBUser = dbUser
	} else {
		return nil, models.ErrBadSSLMode
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	} else {
		return nil, models.ErrBadSSLMode
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	} else {
		return nil, models.ErrBadSSLMode
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	} else {
		return nil, models.ErrBadSSLMode
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	} else {
		return nil, models.ErrBadDBName
	}

	if dbSSLMode := os.Getenv("SSL_MODE"); dbSSLMode != "" {
		config.DBSSLMode = dbSSLMode
	} else {
		return nil, models.ErrBadSSLMode
	}

	if migrationsPath := os.Getenv("MIGRATIONS_PATH"); migrationsPath != "" {
		config.MigrationsPath = migrationsPath
	} else {
		return nil, models.ErrMigrationsNotProvided
	}

	return config, nil
}

func MigratorConnStr(cfg *MigratorConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
