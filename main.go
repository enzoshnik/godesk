package main

import (
	"github.com/gin-gonic/gin"
	"helpdesk/config"
	"helpdesk/routes"
	"helpdesk/utils"
)

func main() {
	// Инициализация базы данных
	utils.LoadEnv()
	config.InitDatabase()

	// Создаем сервер
	router := gin.Default()

	// Подключаем маршруты
	routes.RegisterAuthRoutes(router)
	routes.RegisterTicketRoutes(router)

	// Запускаем сервер
	router.Run(":8080")
}
