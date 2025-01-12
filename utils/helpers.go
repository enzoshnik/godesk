package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func TransformList[T any, R any](input []T, transform func(T) R) []R {
	var result []R
	for _, item := range input {
		result = append(result, transform(item))
	}
	return result
}

// parsePagination — вспомогательная функция для обработки параметров пагинации
func ParsePagination(c *gin.Context) (page int, limit int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	return page, limit
}
