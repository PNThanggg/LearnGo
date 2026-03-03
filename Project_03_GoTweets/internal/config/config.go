package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	Port         string
	JwtSecret    string
	PostgresUser string
	PostgresPass string
	PostgresDb   string
}

func LoadConfig() (*Config, error) {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	var config = &Config{
		DatabaseURL:  os.Getenv("DB_URL"),
		Port:         os.Getenv("PORT"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDb:   os.Getenv("POSTGRES_DB"),
	}
	return config, nil
}
