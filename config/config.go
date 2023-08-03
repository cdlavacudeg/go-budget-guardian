package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	DB  DatabaseConfig
}

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

var config *Config

func LoadConfig() (*Config, error) {
	if config != nil {
		// If the config is already loaded, return it
		return config, nil
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config = &Config{
		Env: os.Getenv("GO_ENV"),
		DB: DatabaseConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}

	return config, nil
}

func GetConfig() (*Config, error) {
	if config == nil {
		_, err := LoadConfig()
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
