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
	Order           Order           `json:"order"`
	TransactionType TransactionType `json:"transaction_type"`
	CreatedAt       time.Time       `json:"created_at" gorm:"default:now();"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"default:now();"`
}
