package bootstrap

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig `json:"server"`
	DbConfig DbConfig     `json:"db"`
}

type ServerConfig struct {
	PORT   int    `json:"port"`
	HOST   string `json:"host"`
	NAME   string `json:"name"`
	SECRET string // from env
	ENV    string // from env
	LEVEL  string // from env
}

type DbConfig struct {
}

func loadFromJSON(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("invalid json in %s: %w", path, err)
	}

	return nil
}

func loadFromENV(cfg *Config, envType string) error {
	if envType == "prod" {
		if err := godotenv.Load("config/.prod.env"); err != nil {
			return err
		}
	} else if envType == "dev" {
		if err := godotenv.Load("config/.dev.env"); err != nil {
			return err
		}
	} else {
	}

	cfg.Server.SECRET = os.Getenv("SERVER_SECRET")
	cfg.Server.ENV = os.Getenv("SERVER_ENV")
	cfg.Server.LEVEL = os.Getenv("SERVER_LEVEL")
	return nil
}

func NewConfig(jsonPath string, envType string) (*Config, error) {
	// Load config

	cfg := &Config{}

	if jsonPath != "" {
		if err := loadFromJSON(jsonPath, cfg); err != nil {
			return nil, err
		}
	}

	err := loadFromENV(cfg, envType)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
