package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}

}

// GetSecretKey возвращает значение SECRET_KEY из переменных окружения
func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set in the environment")
	}
	return secretKey
}
