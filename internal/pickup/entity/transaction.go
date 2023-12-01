package entity

import "time"

type TransactionType string

const (
	TransactionTypePickup  TransactionType = "pickup"
	TransactionTypeReceive TransactionType = "receive"
	TransactionTypeRefund  TransactionType = "refund"
)

type Transaction struct {
	ID              int             `json:"id" gorm:"primary_key"`
	UserID          int             `json:"user_id"`
	OrderID         int             `json:"order_id"`
	TransactionType TransactionType `json:"transaction_type"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Order           Order           `json:"order"`
}
