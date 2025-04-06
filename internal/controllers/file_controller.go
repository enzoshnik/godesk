package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"helpdesk/config"
	"helpdesk/internal/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// UploadFile загружает файл для тикета
func UploadFile(c *gin.Context) {
	ticketID := c.Param("ticket_id") // ID тикета из маршрута

	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл"})
		return
	}

	// Создаем директорию для файла
	uploadPath := filepath.Join("uploads", ticketID)
	fmt.Println(uploadPath)
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать директорию"})
		return
	}

	// Сохраняем файл
	filePath := filepath.Join(uploadPath, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения файла"})
		return
	}

	// Сохраняем информацию о файле в базе данных

	intTicketID, err := strconv.ParseUint(ticketID, 10, 64)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	newFile := models.File{
		TicketID: uint(intTicketID),
		FileName: file.Filename,
		FilePath: filePath,
	}
	if err := config.DB.Create(&newFile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения данных о файле"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Файл успешно загружен", "file": newFile})
}

// DownloadFile возвращает файл по ID
func DownloadFile(c *gin.Context) {
	fileID := c.Param("file_id")

	// Получаем файл из базы данных
	db := c.MustGet("db").(*gorm.DB)
	var file models.File
	if err := db.First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Файл не найден"})
		return
	}

	// Отправляем файл клиенту
	c.File(file.FilePath)
}
