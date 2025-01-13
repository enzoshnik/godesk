package models

import "time"

type Comment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TicketID    uint      `json:"ticket_id" gorm:"not null"` // Внешний ключ к тикету
	Ticket      Ticket    `json:"-" gorm:"foreignKey:TicketID;references:ID"`
	Text        string    `json:"text" gorm:"type:text;not null"` // Текст комментария
	CreatedById uint      `json:"created_by_id" gorm:"not null"`  // Имя автора комментария
	CreatedBy   User      `json:"-" gorm:"foreignKey:CreatedById;references:ID"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;autoCreateTime"`    // Время создания
	ChangeAt    time.Time `json:"change_at" gorm:"not null;autoCreateTime"`     // Время создания
	ChangeById  uint      `json:"change_by_id"`                                 // Кем изменено
	ChangeBy    User      `json:"-" gorm:"foreignKey:ChangeById;references:ID"` // Кем изменено
	Public      bool      `json:"-"`
}
