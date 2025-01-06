package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

// Структура для тикетов
type Ticket struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

// Глобальная переменная для базы данных
var db *gorm.DB

func initDatabase() {
	// Подключение к PostgreSQL
	dsn := "host=localhost user=postgres password=yourpassword dbname=helpdesk port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Миграция структуры
	err = db.AutoMigrate(&Ticket{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

func main() {
	// Инициализация базы данных
	initDatabase()

	router := gin.Default()

	// Маршрут: Получение всех тикетов
	router.GET("/tickets", func(c *gin.Context) {
		var tickets []Ticket
		db.Find(&tickets)
		c.JSON(http.StatusOK, tickets)
	})

	// Маршрут: Создание нового тикета
	router.POST("/tickets", func(c *gin.Context) {
		var newTicket Ticket
		if err := c.ShouldBindJSON(&newTicket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newTicket)
		c.JSON(http.StatusCreated, newTicket)
	})

	// Маршрут: Получение тикета по ID
	router.GET("/tickets/:id", func(c *gin.Context) {
		id := c.Param("id")
		var ticket Ticket
		if err := db.First(&ticket, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
			return
		}
		c.JSON(http.StatusOK, ticket)
	})

	// Маршрут: Обновление тикета
	router.PUT("/tickets/:id", func(c *gin.Context) {
		id := c.Param("id")
		var ticket Ticket
		if err := db.First(&ticket, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
			return
		}
		if err := c.ShouldBindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&ticket)
		c.JSON(http.StatusOK, ticket)
	})

	// Маршрут: Удаление тикета
	router.DELETE("/tickets/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Ticket{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
	})

	router.Run(":8080")
}
