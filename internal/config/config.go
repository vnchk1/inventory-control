package config

import (
	"fmt"
	"os"

	"github.com/vnchk1/inventory-control/internal/models"
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
		return nil, models.ErrBadLogLevel
	}

	if serverPort := os.Getenv("SERVER_PORT"); serverPort != "" {
		config.ServerPort = serverPort
	} else {
		return nil, models.ErrBadPort
	}

	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.DBUser = dbUser
	} else {
		return nil, models.ErrBadUserName
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	} else {
		return nil, models.ErrBadPassword
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	} else {
		return nil, models.ErrBadHost
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	} else {
		return nil, models.ErrBadPort
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

	return config, nil
}

func ConnStr(cfg *Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
