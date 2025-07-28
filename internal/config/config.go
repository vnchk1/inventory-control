package config

import (
	"errors"
	"fmt"
	"os"
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

	if logLvl := os.Getenv("LOG_LEVEL"); logLvl != "" {
		config.LogLevel = logLvl
	} else {
		return &Config{}, errors.New("LOG_LEVEL is required")
	}

	if serverPort := os.Getenv("SERVER_PORT"); serverPort != "" {
		config.ServerPort = serverPort
	} else {
		return &Config{}, errors.New("SERVER_PORT is required")
	}

	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.DBUser = dbUser
	} else {
		return &Config{}, errors.New("DB_USER is required")
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	} else {
		return &Config{}, errors.New("DB_PASSWORD is required")
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	} else {
		return &Config{}, errors.New("DB_HOST is required")
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	} else {
		return &Config{}, errors.New("DB_PORT is required")
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	} else {
		return &Config{}, errors.New("DB_NAME is required")
	}

	if dbSSLMode := os.Getenv("SSL_MODE"); dbSSLMode != "" {
		config.DBSSLMode = dbSSLMode
	} else {
		return &Config{}, errors.New("SSL_MODE is required")
	}

	return config, nil
}

func ConnStr(cfg *Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
