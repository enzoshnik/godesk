package controllers

import (
	"github.com/gin-gonic/gin"
	"helpdesk/config"
	"helpdesk/models"
	"helpdesk/utils"
	"net/http"
)

func Tickets(context *gin.Context) {
	var tickets []models.Ticket
	config.DB.Find(&tickets)
	context.JSON(http.StatusOK, tickets)
}

func CreateTicket(context *gin.Context) {
	var ticket models.Ticket
	if err := context.ShouldBindJSON(&ticket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket.CreatedBy = context.GetString("username")
	config.DB.Create(&ticket)
	context.JSON(http.StatusCreated, ticket)
}

func MyTickets(context *gin.Context) {
	username := context.GetString("username")
	var tickets []models.Ticket
	config.DB.Preload("Status").Where("created_by = ?", username).Find(&tickets)

	// Преобразование списка с помощью универсальной функции
	transformedTickets := utils.TransformList(tickets, func(ticket models.Ticket) models.TicketFotList {
		return models.TicketFotList{
			ID:        ticket.ID,
			Title:     ticket.Title,
			Content:   ticket.Content,
			Status:    ticket.Status,
			CreatedBy: ticket.CreatedBy,
		}
	})

	context.JSON(http.StatusOK, transformedTickets)
}

func DeleteTicket(c *gin.Context) {
	// Получаем ID тикета из параметров запроса
	ticketID := c.Param("id")
	role := c.GetString("role")

	// Проверка прав доступа
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied. Only administrators can delete tickets."})
		return
	}

	// Проверяем, существует ли тикет
	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
		return
	}

	// Удаление тикета
	if err := config.DB.Delete(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
