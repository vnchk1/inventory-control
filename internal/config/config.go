package config

import (
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
	}
	if serverPort := os.Getenv("SERVER_PORT"); serverPort != "" {
		config.ServerPort = serverPort
	}
	if user := os.Getenv("DB_USER"); user != "" {
		config.DBUser = user
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		config.DBPassword = password
	}
	if host := os.Getenv("DB_HOST"); host != "" {
		config.DBHost = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		config.DBPort = port
	}
	if dbname := os.Getenv("DB_NAME"); dbname != "" {
		config.DBName = dbname
	}
	if sslmode := os.Getenv("SSL_MODE"); sslmode != "" {
		config.DBSSLMode = sslmode
	}

	return config, nil
}
