package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
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

var (
	ENV      = getENV("ENV", "testing")
	AppName  = "sea-labs-library"
	DBConfig = dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", "postgres"),
		Password: getENV("DB_PASSWORD", "postgres"),
		DBName:   getENV("DB_NAME", "library_db"),
		Port:     getENV("DB_PORT", "5432"),
	}
)
