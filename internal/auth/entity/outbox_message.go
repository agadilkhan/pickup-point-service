package entity

type OutboxMessage struct {
	ID          int    `json:"id" gorm:"primary_key"`
	UserEmail   string `json:"user_email"`
	Code        string `json:"code"`
	IsProcessed bool   `json:"is_processed"`
}
