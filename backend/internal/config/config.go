package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	DB_URL string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		PORT: os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}, nil
}
