package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrBadHost       = errors.New("host is required")
	ErrBadServerPort = errors.New("server port is required")
	ErrBadDBPort     = errors.New("db port is required")
	ErrBadSSLMode    = errors.New("SSL_MODE is required")
	ErrBadUserName   = errors.New("username is required")
	ErrBadPassword   = errors.New("password is required")
	ErrBadDBName     = errors.New("db name is required")
	ErrBadLogLevel   = errors.New("log level is required")
)

type Config struct {
	LogLevel   string
	ServerPort string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	envConfigs := map[string]*string{
		"LOG_LEVEL":   &config.LogLevel,
		"SERVER_PORT": &config.ServerPort,
		"DB_USER":     &config.DBUser,
		"DB_PASSWORD": &config.DBPassword,
		"DB_HOST":     &config.DBHost,
		"DB_PORT":     &config.DBPort,
		"DB_NAME":     &config.DBName,
		"SSL_MODE":    &config.DBSSLMode,
	}

	envErrors := map[string]error{
		"LOG_LEVEL":   ErrBadLogLevel,
		"SERVER_PORT": ErrBadServerPort,
		"DB_USER":     ErrBadUserName,
		"DB_PASSWORD": ErrBadPassword,
		"DB_HOST":     ErrBadHost,
		"DB_PORT":     ErrBadDBPort,
		"DB_NAME":     ErrBadDBName,
		"SSL_MODE":    ErrBadSSLMode,
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

func ConnStr(cfg *Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
