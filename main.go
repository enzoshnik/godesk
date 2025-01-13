package main

import (
	"github.com/gin-gonic/gin"
	"helpdesk/config"
	"helpdesk/docs"
	"helpdesk/routes"
	"helpdesk/utils"
)

func main() {
	// Инициализация базы данных
	utils.LoadEnv()
	config.InitDatabase()

	// Создаем сервер
	router := gin.Default()

	// Swagger настройка
	docs.SwaggerInfo.Title = "Helpdesk API"
	docs.SwaggerInfo.Description = "API documentation for the Helpdesk system"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	// Подключаем маршруты
	routes.RegisterAuthRoutes(router)
	routes.RegisterTicketRoutes(router)
	routes.RegisterSwaggerRoutes(router)

	// Запускаем сервер
	router.Run(":8080")
}
