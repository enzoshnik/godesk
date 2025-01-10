package models

import "encoding/json"

type TicketFotList struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StatusID  uint   `json:"status_id"`
	Status    Status `json:"-"`
	CreatedBy string `json:"created_by"`
}

// Кастомный метод MarshalJSON для управления выводом Status в списке Tickets
func (ct TicketFotList) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID         uint   `json:"id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		Status     Status `json:"-"`
		StatusID   uint   `json:"-"`
		StatusCode string `json:"status_code"`
		CreatedBy  string `json:"created_by"`
	}{
		ID:         ct.ID,
		Title:      ct.Title,
		Content:    ct.Content,
		StatusID:   ct.StatusID,
		StatusCode: ct.Status.Code, // Только code
		CreatedBy:  ct.CreatedBy,
	})
}
