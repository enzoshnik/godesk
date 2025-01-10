package controllers

import (
	"github.com/gin-gonic/gin"
	"helpdesk/config"
	"helpdesk/models"
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
	config.DB.Where("created_by = ?", username).Find(&tickets)
	context.JSON(http.StatusOK, tickets)
}
