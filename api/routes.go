package api

import (
	"github.com/gin-gonic/gin"
	"helpdesk/api/swagger"
	"helpdesk/api/v1"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	swagger.RegisterSwaggerRoutes(router)

	// Версия 1
	apiPath := api.Group("/v1")
	v1.RegisterAuthRoutes(apiPath)
	v1.RegisterTicketRoutes(apiPath)
}
