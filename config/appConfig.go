package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (AppConfig, error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	dsn := os.Getenv("DB_SOURCE")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{
			ServerPort: httpPort,
			Dsn:        dsn,
		},
		nil
}
