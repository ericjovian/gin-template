package config

import (
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	ENV = getENV("ENV", "testing")
	AppName = "gin-template-library"
	DBConfig = dbConfig{
		Host:     getENV("DB_HOST"),
		User:     getENV("DB_USER"),
		Password: getENV("DB_PASSWORD"),
		DBName:   getENV("DB_NAME"),
		Port:     getENV("DB_PORT"),
	}
}
