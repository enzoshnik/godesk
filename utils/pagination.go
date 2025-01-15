package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type PaginationResult[T any] struct {
	Data       []T   `json:"data"`        // Список элементов
	TotalCount int64 `json:"total_count"` // Общее количество элементов
	Page       int   `json:"page"`        // Текущая страница
	PageSize   int   `json:"page_size"`   // Размер страницы
}

// Paginate обрабатывает параметры пагинации и возвращает результат.
func Paginate[T any](db *gorm.DB, c *gin.Context, model *[]T) (PaginationResult[T], error) {
	// Получение параметров из запроса
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Вычисление пропускаемых записей
	offset := (page - 1) * pageSize

	// Получение данных и общего количества
	var totalCount int64
	if err := db.Model(model).Count(&totalCount).Error; err != nil {
		return PaginationResult[T]{}, err
	}

	if err := db.Limit(pageSize).Offset(offset).Find(model).Error; err != nil {
		return PaginationResult[T]{}, err
	}

	// Формирование результата
	return PaginationResult[T]{
		Data:       *model,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}
