package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		PORT: os.Getenv("PORT"),
	}, nil
}
