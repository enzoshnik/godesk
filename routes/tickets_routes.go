package routes

import (
	"github.com/gin-gonic/gin"
	"helpdesk/controllers"
	"helpdesk/middlewares"
)

func RegisterTicketRoutes(router *gin.Engine) {
	protected := router.Group("/tickets")
	protected.Use(middlewares.AuthenticateMiddleware())
	// Только администратор: получить все тикеты
	protected.GET("/", middlewares.AdminOnlyMiddleware(), controllers.Tickets)
	// Пользователь: создать тикет
	protected.POST("/", controllers.CreateTicket)
	// Пользователь: просмотреть свои тикеты
	protected.GET("/my", controllers.MyTickets)
	protected.PATCH("/:id/status", controllers.UpdateTicketStatus) // Изменение статуса
	protected.DELETE("/:id", controllers.DeleteTicket)             // Администратор: удалить тикет

	// comment
	protected.POST("/comments", controllers.AddComment)               // Добавление комментария
	protected.GET("/:id/comments", controllers.GetCommentsByTicketID) // Получение комментариев

}
