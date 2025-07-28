package config

import (
	"errors"
	"fmt"
	"os"
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
		return &MigratorConfig{}, errors.New("DB_USER is required")
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	} else {
		return &MigratorConfig{}, errors.New("DB_PASSWORD is required")
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	} else {
		return &MigratorConfig{}, errors.New("DB_HOST is required")
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	} else {
		return &MigratorConfig{}, errors.New("DB_PORT is required")
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	} else {
		return &MigratorConfig{}, errors.New("DB_NAME is required")
	}

	if dbSSLMode := os.Getenv("SSL_MODE"); dbSSLMode != "" {
		config.DBSSLMode = dbSSLMode
	} else {
		return &MigratorConfig{}, errors.New("DB_SSL_MODE is required")
	}

	if migrationsPath := os.Getenv("MIGRATIONS_PATH"); migrationsPath != "" {
		config.MigrationsPath = migrationsPath
	} else {
		return &MigratorConfig{}, errors.New("MIGRATIONS_PATH is required")
	}

	return config, nil
}

func MigratorConnStr(cfg *MigratorConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
