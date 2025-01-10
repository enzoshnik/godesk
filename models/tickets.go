package models

type Ticket struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StatusId  int    `json:"status_id" gorm:"default:1"`        // Статусы: "Открыт", "В работе", "Закрыт"
	Status    Status `json:"status" gorm:"foreignKey:StatusId"` // Reference
	CreatedBy string `json:"created_by"`
}
