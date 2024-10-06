package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }
}

func GetPort() string {
    return os.Getenv("PORT")
}

func GetMongoURI() string {
    return os.Getenv("MONGODB_URI")
}
