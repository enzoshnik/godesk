package models

import "time"

type File struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TicketID  uint      `gorm:"index;not null" json:"ticket_id"` // Связь с тикетом
	FileName  string    `gorm:"not null" json:"file_name"`
	FilePath  string    `gorm:"not null" json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
}
