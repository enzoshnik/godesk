package controllers

import (
	"github.com/gin-gonic/gin"
	"helpdesk/config"
	"helpdesk/models"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	var comment models.Comment

	// Привязка данных из запроса
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ChangeById = comment.CreatedById

	// Проверка существования тикета
	var ticket models.Ticket
	if err := config.DB.First(&ticket, comment.TicketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
		return
	}

	// Сохранение комментария
	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully", "comment": comment})
}

// GetCommentsByTicketID retrieves all comments for a specific ticket
func GetCommentsByTicketID(c *gin.Context) {
	// Получаем TicketID из параметров URL
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ticket ID"})
		return
	}

	// Проверяем, существует ли тикет
	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
		return
	}

	// Получаем параметр сортировки (по умолчанию "asc")
	sort := c.DefaultQuery("sort", "asc")
	if sort != "asc" && sort != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid sort parameter. Use 'asc' or 'desc'"})
		return
	}

	// Извлекаем комментарии для тикета
	var comments []models.Comment
	order := "created_at " + sort
	if err := config.DB.Where("ticket_id = ?", ticketID).Order(order).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ticket_id": ticketID,
		"comments":  comments,
	})
}
