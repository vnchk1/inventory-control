package config

import (
	"errors"
	"fmt"
	"os"
)

var ErrMigrationsNotProvided = errors.New("migrations not provided")

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
		return nil, ErrBadSSLMode
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	} else {
		return nil, ErrBadSSLMode
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	} else {
		return nil, ErrBadSSLMode
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	} else {
		return nil, ErrBadSSLMode
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	} else {
		return nil, ErrBadDBName
	}

	if dbSSLMode := os.Getenv("SSL_MODE"); dbSSLMode != "" {
		config.DBSSLMode = dbSSLMode
	} else {
		return nil, ErrBadSSLMode
	}

	if migrationsPath := os.Getenv("MIGRATIONS_PATH"); migrationsPath != "" {
		config.MigrationsPath = migrationsPath
	} else {
		return nil, ErrMigrationsNotProvided
	}

	return config, nil
}

func MigratorConnStr(cfg *MigratorConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
