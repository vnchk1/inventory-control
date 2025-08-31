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

	envConfigs := map[string]*string{
		"DB_USER":         &config.DBUser,
		"DB_PASSWORD":     &config.DBPassword,
		"DB_HOST":         &config.DBHost,
		"DB_PORT":         &config.DBPort,
		"DB_NAME":         &config.DBName,
		"SSL_MODE":        &config.DBSSLMode,
		"MIGRATIONS_PATH": &config.MigrationsPath,
	}

	envErrors := map[string]error{
		"DB_USER":         ErrBadUserName,
		"DB_PASSWORD":     ErrBadPassword,
		"DB_HOST":         ErrBadHost,
		"DB_PORT":         ErrBadDBPort,
		"DB_NAME":         ErrBadDBName,
		"SSL_MODE":        ErrBadSSLMode,
		"MIGRATIONS_PATH": ErrMigrationsNotProvided,
	}

	for envName, configField := range envConfigs {
		value := os.Getenv(envName)
		if value == "" {
			return nil, envErrors[envName]
		}

		*configField = value
	}

	return config, nil
}

func MigratorConnStr(cfg *MigratorConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
