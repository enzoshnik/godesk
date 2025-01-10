package controllers

import (
	"fmt"
	"helpdesk/config"
	"helpdesk/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo полностью переделать
func UpdateTicketStatus(c *gin.Context) {
	// Получение ID тикета из параметров запроса
	ticketID := c.Param("id")
	username := c.GetString("username")
	role := c.GetString("role")

	// Проверяем, существует ли тикет
	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
		return
	}

	// Проверка прав доступа
	if role != "admin" && ticket.CreatedBy != username {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return
	}

	// Получение нового статуса из тела запроса
	var request struct {
		Status string `json:"status"`
	}

	var status models.Status
	if err := config.DB.Where(&models.Status{Code: request.Status}).First(&status).Error; err != nil {
		fmt.Println("Error querying users:", err)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Валидация статуса
	allowedStatuses := map[string]bool{"open": true, "in_progress": true, "closed": true}
	if !allowedStatuses[status.Code] {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid status"})
		return
	}

	// Обновление статуса
	ticket.StatusId = status.ID
	if err := config.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket status updated", "ticket": ticket})
}
