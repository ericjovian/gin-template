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

func getENV(key string) string {
	return os.Getenv(key)
}

var (
	ENV      = getENV("ENV")
	AppName  = "gin-template-library"
	DBConfig dbConfig
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	DBConfig = dbConfig{
		Host:     getENV("DB_HOST"),
		User:     getENV("DB_USER"),
		Password: getENV("DB_PASSWORD"),
		DBName:   getENV("DB_NAME"),
		Port:     getENV("DB_PORT"),
	}
}
