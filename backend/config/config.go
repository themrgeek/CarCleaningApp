package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	AppPort    string
	LogFile    string
}

func LoadConfig() Config {

	appPort, exists := os.LookupEnv("APP_PORT")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if !exists {

		log.Fatal("APP_PORT environment variable not set")

	}

	logFile, exists := os.LookupEnv("LOG_FILE")

	if !exists {

		log.Fatal("LOG_FILE environment variable not set")

	}

	return Config{

		AppPort: appPort,

		LogFile:    logFile,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		// Initialize other fields as needed

	}

}
