package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	serverAddr string
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	dbUser = getEnv("DB_USER", "root")
	dbPassword = getEnv("DB_PASSWORD", "password")
	dbHost = getEnv("DB_HOST", "localhost")
	dbName = getEnv("DB_NAME", "lms")
	dbPort = getEnv("DB_PORT", "3306")
	serverAddr = getEnv("SERVER_ADDR", ":8080")

}

func GetDBConnectionString() string {
	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
}

func GetServerAddress() string {
	return serverAddr
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
