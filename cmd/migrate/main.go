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
	cfg.InitDatabase()

	// Миграция структуры
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Ticket{},
		&models.StatusRelation{},
		&models.Status{},
		&models.Comment{},
		&models.File{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	log.Println("Миграции успешно выполнены")
}
