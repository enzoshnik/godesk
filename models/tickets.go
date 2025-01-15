package models

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	Title                string    `gorm:"size:255;not null"`
	Content              string    `json:"content"`
	StatusId             int       `json:"status_id" gorm:"default:1"`
	Status               Status    `json:"-" gorm:"foreignKey:StatusId;references:ID"` // Reference
	Uuid                 string    `gorm:"size:255;not null"`
	DateCreated          time.Time `gorm:"not null;autoCreateTime"`
	CreatedBy            uint      `gorm:"not null"`
	DateChanged          time.Time `gorm:"not null"`
	ChangedBy            uint      `gorm:"not null"`
	StatusChangedDate    time.Time `gorm:"not null"`
	Description          string    `gorm:"type:text"`
	PriorityId           uint      `gorm:"not null"`
	TypeId               uint      `gorm:"not null"`
	CompanyId            uint      `gorm:"not null"`
	ClientId             uint      `gorm:"not null"`
	ObjectServiceId      uint      `gorm:"not null"`
	ContractId           uint      `gorm:"not null"`
	ResponsibleManagerId uint      `gorm:"not null"`
	Deadline             time.Time `gorm:"not null"`
	Reaction             time.Time `gorm:"not null"`
	NoAnswer             string    `gorm:"size:1;not null"`
	DateClosed           time.Time `gorm:"not null"`
	DateFirstReaction    time.Time `gorm:"not null"`
	DurationPlan         uint      `gorm:"not null"`
	SectionId            uint      `gorm:"not null"`
	ExtFields            string    `gorm:"type:text"`
	Comments             []Comment `json:"-" gorm:"foreignKey:TicketID;references:ID"` // Связь с комментариями
}

func (T *Ticket) BeforeSave(tx *gorm.DB) (err error) {
	now := time.Now()
	T.Deadline = now.Add(3 * 24 * time.Hour)
	return
}
