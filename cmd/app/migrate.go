package main

import (
	"helpdesk/config"
	"helpdesk/internal/models"
	"log"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загружаем конфигурацию (например, из .env)
	cfg := config.LoadConfig()
	db := cfg.InitDatabase()

	// Миграция структуры
	err := db.AutoMigrate(
		&models.User{},
		&models.Ticket{},
		&models.StatusRelation{},
		&models.Status{},
		&models.Comment{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	log.Println("Миграции успешно выполнены")
}
