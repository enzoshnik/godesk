package models

type Ticket struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status" gorm:"default:'Открыт'"` // Статусы: "Открыт", "В работе", "Закрыт"
	CreatedBy string `json:"created_by"`
}
