package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
)

type Config struct {
	Env          string `env:"ENV"`
	DATABASE_URL string `env:"DATABASE_URL"`
}

func NewConfig() *Config {
	var config Config
	switch os.Getenv("ENV") {
	case EnvDevelopment:
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		if err := env.Parse(&config); err != nil {
			log.Fatal(err)
		}

	default:
		panic("Invalid ENV")
	}

	return &config
}
