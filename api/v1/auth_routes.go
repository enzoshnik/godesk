package v1

import (
	"github.com/gin-gonic/gin"
	"helpdesk/internal/controllers"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
