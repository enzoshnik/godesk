package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"helpdesk/models"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=helpdesk port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Миграция структуры
	err = DB.AutoMigrate(
		&models.User{},
		&models.Ticket{},
		&models.StatusRelation{},
		&models.Status{},
		&models.Comment{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	(&models.Status{}).Install(DB)
}
