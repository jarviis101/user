package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDSN string
	Debug       bool
	Secret      string
	Host        string
	Port        string
}

func LoadConfig() (*Config, error) {
	if err := loadEnv(); err != nil {
		return nil, err
	}

	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))

	if err != nil {
		debug = false
	}

	fmt.Print(os.Getenv("APP_DEBUG"))

	return &Config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
		Secret:      os.Getenv("APP_KEY"),
		Host:        os.Getenv("APP_HOST"),
		Port:        os.Getenv("APP_PORT"),
		Debug:       debug,
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
