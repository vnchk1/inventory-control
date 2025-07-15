package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	LogLevel string
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file %w", err)
	}

	if logLvl := os.Getenv("LOG_LEVEL"); logLvl != "" {
		config.LogLevel = logLvl
	}
	if user := os.Getenv("USER"); user != "" {
		config.User = user
	}
	if password := os.Getenv("PASSWORD"); password != "" {
		config.Password = password
	}
	if host := os.Getenv("HOST"); host != "" {
		config.Host = host
	}
	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	}
	if dbname := os.Getenv("DB_NAME"); dbname != "" {
		config.DBName = dbname
	}
	if sslmode := os.Getenv("SSL_MODE"); sslmode != "" {
		config.SSLMode = sslmode
	}

	return config, nil
}
