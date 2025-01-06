package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер Gin
	router := gin.Default()

	// Настраиваем базовый маршрут для проверки работы сервера
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Helpdesk System",
		})
	})

	// Запускаем сервер на порту 8080
	router.Run(":8080")
}
