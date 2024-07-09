package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PublicHost string
	Port       string
	Mode       string
	Database   Database
}

var Env = initConfig()

func initConfig() config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		Mode:       getEnv("MODE", "development"),
		Database: Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DbName:   getEnv("DB_NAME", "postgres"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
