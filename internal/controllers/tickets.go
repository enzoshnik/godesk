package controllers

import (
	"fmt"
	"helpdesk/config"
	"helpdesk/internal/models"
	utils2 "helpdesk/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllTickets godoc
// @Summary Get all tickets
// @Description Retrieve a list of tickets with optional filters and pagination
// @Tags Tickets
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page number"
// @Param limit query int false "Limit of tickets per page"
// @Param status query string false "Filter by ticket status"
// @Param created_by query string false "Filter by creator username"
// @Param start_date query string false "Filter by start date (YYYY-MM-DD)"
// @Param end_date query string false "Filter by end date (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/v1/tickets/ [get]
func Tickets(c *gin.Context) {
	// Параметры пагинации
	page, limit := utils2.ParsePagination(c)
	var status models.Status

	// Параметры фильтрации
	statusCode := c.Query("status")
	createdBy := c.Query("created_by")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// Формируем запрос
	var tickets []models.Ticket
	query := config.DB.Model(&models.Ticket{})

	// Применяем фильтры
	if statusCode != "" {
		// Находим статус по коду
		if err := config.DB.First(&status, "code = ?", statusCode).Error; err != nil {
			fmt.Println("Status not found:", err)
			return
		}

		query = query.Where("status_id = ?", status.ID)
	}
	if createdBy != "" {
		query = query.Where("created_by = ?", createdBy)
	}
	if startDate != "" {
		start, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}
		query = query.Where("created_at >= ?", start)
	}
	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}
		query = query.Where("created_at <= ?", end)
	}

	// todo Добавить фильтр по ответственному

	// Пагинация
	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	// Выполняем запрос
	if err := query.Find(&tickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch tickets"})
		return
	}

	// Считаем общее количество тикетов (с учётом фильтров)
	var total int64
	query.Count(&total)

	result, err := utils2.Paginate(config.DB, c, &tickets)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve tickets", "error": err.Error()})
		return
	}

	// Формируем ответ
	c.JSON(http.StatusOK, result)

}

func CreateTicket(context *gin.Context) {
	var ticket models.Ticket
	if err := context.ShouldBindJSON(&ticket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUUID := uuid.New()
	ticket.Uuid = newUUID.String()

	ticket.CreatedBy = context.GetUint("userID")
	config.DB.Create(&ticket)
	context.JSON(http.StatusCreated, ticket)
}

func MyTickets(context *gin.Context) {

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit value"})
		return
	}

	// Вычисляем смещение (offset)
	offset := (page - 1) * limit

	userId := context.GetUint("userId")
	var tickets []models.Ticket
	config.DB.Preload("Status").Where("created_by = ?", userId).Limit(limit).Offset(offset).Find(&tickets)

	// Преобразование списка с помощью универсальной функции
	transformedTickets := utils2.TransformList(tickets, func(ticket models.Ticket) models.TicketForList {
		return models.TicketForList{
			ID:        ticket.ID,
			Title:     ticket.Title,
			Content:   ticket.Content,
			Status:    ticket.Status,
			CreatedBy: ticket.CreatedBy,
		}
	})

	var total int64
	config.DB.Model(&models.Ticket{}).Count(&total)

	context.JSON(http.StatusOK, gin.H{
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // Округление вверх
		"tickets":    transformedTickets,
	})
}

func DeleteTicket(context *gin.Context) {
	// Получаем ID тикета из параметров запроса
	ticketID := context.Param("id")
	role := context.GetString("role")

	// Проверка прав доступа
	if role != "admin" {
		context.JSON(http.StatusForbidden, gin.H{"message": "Access denied. Only administrators can delete tickets."})
		return
	}

	// Проверяем, существует ли тикет
	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
		return
	}

	// Удаление тикета
	if err := config.DB.Delete(&ticket).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete ticket"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
