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

	if logLvl := os.Getenv("LOG_LEVEL"); logLvl != "" {
		config.LogLevel = logLvl
	} else {
		return nil, ErrBadLogLevel
	}

	if serverPort := os.Getenv("SERVER_PORT"); serverPort != "" {
		config.ServerPort = serverPort
	} else {
		return nil, ErrBadServerPort
	}

	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.DBUser = dbUser
	} else {
		return nil, ErrBadUserName
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	} else {
		return nil, ErrBadPassword
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	} else {
		return nil, ErrBadHost
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	} else {
		return nil, ErrBadDBPort
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

	return config, nil
}

func ConnStr(cfg *Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
}
