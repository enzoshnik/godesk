package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Status struct {
	ID          int              `json:"id" gorm:"primaryKey"`
	Title       string           `json:"title"`
	DateCreated time.Time        `json:"date_created" gorm:"autoCreateTime"`
	Button      string           `json:"button"`
	Code        string           `json:"code"`
	Visibility  string           `json:"visibility"`
	Color       string           `json:"color"`
	Default     bool             `json:"default"`
	FromStatus  []StatusRelation `json:"from_status" gorm:"foreignKey:StatusFromID;references:ID"` // Reference
	ToStatus    []StatusRelation `json:"to_status" gorm:"foreignKey:StatusToID;references:ID"`     // Reference
	Final       bool             `json:"final"`
	Autoclose   int              `json:"autoclose"`
	Status      string           `json:"status" gorm:"default:'Открыт'"` // Статусы: "Открыт", "В работе", "Закрыт"
}

func (Status) Install(db *gorm.DB) {
	// Проверяем, есть ли данные в таблице
	var count int64
	db.Model(&Status{}).Count(&count)

	if count > 0 {
		fmt.Printf("В таблице Status уже есть %d записей.\n", count)
	} else {
		fmt.Println("Таблица Status пуста, добавляем данные...")
		// Заполняем таблицу значениями по умолчанию
		defaultStatuses := []Status{
			{Title: "Открыт", Button: "Начать", Code: "open", Color: "#00FF00", Default: true, Final: false, Autoclose: 0},
			{Title: "В работе", Button: "Закрыть", Code: "in_progress", Color: "#FFFF00", Default: false, Final: false, Autoclose: 0},
			{Title: "Закрыт", Button: "Повторить", Code: "closed", Color: "#FF0000", Default: false, Final: true, Autoclose: 1},
		}
		db.Create(&defaultStatuses)
		fmt.Println("Данные добавлены.")
	}
}

func (u *Status) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (u *Status) BeforeSave(tx *gorm.DB) (err error) {
	return
}

// AfterSave /**
func (u *Status) AfterSave(tx *gorm.DB) (err error) {
	return
}

func (u *Status) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

type StatusRelation struct {
	ID           int `gorm:"primaryKey"`
	StatusFromID int `gorm:"not null"` // FK: ссылается на Status.ID
	StatusToID   int `gorm:"not null"` // FK: ссылается на Status.ID
}
