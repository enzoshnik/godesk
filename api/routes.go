package api

import (
	"helpdesk/api/swagger"
	v1 "helpdesk/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	swagger.RegisterSwaggerRoutes(router)

	// Версия 1
	apiPath := api.Group("/v1")
	v1.RegisterAuthRoutes(apiPath)
	v1.RegisterTicketRoutes(apiPath)
	v1.RegisterFileRoutes(apiPath)
}
