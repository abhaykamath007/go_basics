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
	jwtSecret  string
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	} else {
		log.Println(".env file loaded successfully")
	}

	dbUser = getEnv("DB_USER", "root")
	dbPassword = getEnv("DB_PASSWORD", "mysql123#")
	dbHost = getEnv("DB_HOST", "localhost")
	dbName = getEnv("DB_NAME", "lms")
	dbPort = getEnv("DB_PORT", "3306")
	serverAddr = getEnv("SERVER_ADDR", ":8080")
	jwtSecret = getEnv("JWT_SECRET", "p3cr3tK3y@123!yH1sIsA$3cur3K3y")

	log.Println("DB_USER:", dbUser)
	log.Println("DB_PASSWORD:", dbPassword)
	log.Println("DB_HOST:", dbHost)
	log.Println("DB_PORT:", dbPort)
	log.Println("DB_NAME:", dbName)
	log.Println("SERVER_ADDR:", serverAddr)
	log.Println("JWT_SECRET:", jwtSecret)
}

func GetDBConnectionString() string {
	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
}

func GetServerAddress() string {
	return serverAddr
}

func GetSecretKey() string {
	log.Println(jwtSecret)
	return jwtSecret
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
