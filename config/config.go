package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type AuthConfig struct {
	SecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
		Auth: AuthConfig{
			SecretKey: os.Getenv("SECRET_KEY"),
		},
	}
}

func (config *Config) InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Db.Host, config.Db.User, config.Db.Password, config.Db.Name, config.Db.Port)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return DB
}
