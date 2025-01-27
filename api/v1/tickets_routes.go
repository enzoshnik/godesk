package v1

import (
	"github.com/gin-gonic/gin"
	controllers2 "helpdesk/internal/controllers"
	"helpdesk/internal/middlewares"
)

func RegisterTicketRoutes(router *gin.RouterGroup) {
	protected := router.Group("/tickets")
	protected.Use(middlewares.AuthenticateMiddleware())
	// Только администратор: получить все тикеты
	protected.GET("/", middlewares.AdminOnlyMiddleware(), controllers2.Tickets)
	// Пользователь: создать тикет
	protected.POST("/", controllers2.CreateTicket)
	// Пользователь: просмотреть свои тикеты
	protected.GET("/my", controllers2.MyTickets)
	protected.PATCH("/:id/status", controllers2.UpdateTicketStatus) // Изменение статуса
	protected.DELETE("/:id", controllers2.DeleteTicket)             // Администратор: удалить тикет

	// comment
	protected.POST("/comments", controllers2.AddComment)               // Добавление комментария
	protected.GET("/:id/comments", controllers2.GetCommentsByTicketID) // Получение комментариев

}
