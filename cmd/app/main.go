package main

import (
	"fmt"
	"helpdesk/api"
	"helpdesk/api/docs"
	"helpdesk/config"

	// "helpdesk/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	t := time.Now()
	// Инициализация базы данных
	// utils.LoadEnv()
	conf := config.LoadConfig()
	conf.InitDatabase()

	// Создаем сервер
	router := gin.Default()

	// Swagger настройка
	docs.SwaggerInfo.Title = "Helpdesk API"
	docs.SwaggerInfo.Description = "API documentation for the Helpdesk system"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	// Подключаем маршруты
	api.RegisterRoutes(router)

	fmt.Println(time.Since(t))

	// Запускаем сервер
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
