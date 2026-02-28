package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
	Port        string `json:"port"`
}

func LoadConfig() (*Config, error) {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	var config = &Config{
		DatabaseURL: os.Getenv("DB_URL"),
		Port:        os.Getenv("PORT"),
	}
	return config, nil
}
