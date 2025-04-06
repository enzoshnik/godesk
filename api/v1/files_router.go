package v1

import (
	"helpdesk/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(router *gin.RouterGroup) {
	files := router.Group("/files")
	{
		files.POST("/:ticket_id/upload", controllers.UploadFile)  // Загрузка файла
		files.GET("/:file_id/download", controllers.DownloadFile) // Скачивание файла
	}
}
