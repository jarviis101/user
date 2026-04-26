package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDSN string
}

func LoadConfig() (*Config, error) {
	if err := loadEnv(); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
	}, nil
}

func loadEnv() error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	for {
		envPath := filepath.Join(dir, ".env")

		if _, err := os.Stat(envPath); err == nil {
			return godotenv.Load(envPath)
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			break
		}

		dir = parent
	}

	return os.ErrNotExist
}
