package config

import (
	"fmt"
	"github.com/vnchk1/inventory-control/internal/models"
	"os"
)

type Config struct {
	Log      *LoggingConfig
	Server   *ServerConfig
	DB       *DatabaseConfig
	Migrator *MigratorConfig
}

type LoggingConfig struct {
	Level string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type MigratorConfig struct {
	Path string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Log:      &LoggingConfig{},
		Server:   &ServerConfig{},
		DB:       &DatabaseConfig{},
		Migrator: &MigratorConfig{},
	}

	envConfigs := map[string]*string{
		"LOG_LEVEL":       &config.Log.Level,
		"SERVER_PORT":     &config.Server.Port,
		"DB_USER":         &config.DB.Username,
		"DB_PASSWORD":     &config.DB.Password,
		"DB_HOST":         &config.DB.Host,
		"DB_PORT":         &config.DB.Port,
		"DB_NAME":         &config.DB.DBName,
		"SSL_MODE":        &config.DB.SSLMode,
		"MIGRATIONS_PATH": &config.Migrator.Path,
	}

	for envName, configField := range envConfigs {
		value := os.Getenv(envName)
		if value == "" {
			return nil, models.NewEmptyErr(envName)
		}

		*configField = value
	}

	return config, nil
}

func ConnStr(cfg *Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode)
}
