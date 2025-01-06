package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Структура для тикетов
type Ticket struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

// Хранилище тикетов
var tickets = []Ticket{
	{ID: 1, Title: "First Ticket", Content: "This is the first ticket", Status: "open"},
}

// Генерация следующего ID
var nextID = 2

func main() {
	router := gin.Default()

	// Маршрут: Получение всех тикетов
	router.GET("/tickets", func(c *gin.Context) {
		c.JSON(http.StatusOK, tickets)
	})

	// Маршрут: Создание нового тикета
	router.POST("/tickets", func(c *gin.Context) {
		var newTicket Ticket
		if err := c.ShouldBindJSON(&newTicket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTicket.ID = nextID
		nextID++
		tickets = append(tickets, newTicket)
		c.JSON(http.StatusCreated, newTicket)
	})

	// Маршрут: Получение тикета по ID
	router.GET("/tickets/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, ticket := range tickets {
			if c.Param("id") == id {
				c.JSON(http.StatusOK, ticket)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
	})

	// Маршрут: Обновление тикета
	router.PUT("/tickets/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTicket Ticket
		if err := c.ShouldBindJSON(&updatedTicket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, ticket := range tickets {
			if id == string(ticket.ID) {
				tickets[i].Title = updatedTicket.Title
				tickets[i].Content = updatedTicket.Content
				tickets[i].Status = updatedTicket.Status
				c.JSON(http.StatusOK, tickets[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
	})

	// Маршрут: Удаление тикета
	router.DELETE("/tickets/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, ticket := range tickets {
			if id == string(ticket.ID) {
				tickets = append(tickets[:i], tickets[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
	})

	router.Run(":8080")
}
