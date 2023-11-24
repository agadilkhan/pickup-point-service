package entity

import "time"

type PickupOrder struct {
	ID        int       `json:"id" db:"id" gorm:"primary_key;"`
	UserID    int       `json:"user_id" db:"user_id"`
	OrderID   int       `json:"order_id" db:"order_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at" gorm:"default:now();"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" gorm:"default:now();"`
	Order     Order     `json:"order" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type PickupPoint struct {
	ID      int    `json:"id" db:"id" gorm:"primary_key;"`
	Name    string `json:"name" db:"name" gorm:"size:255;"`
	Address string `json:"address" db:"address" gorm:"size:255"`
}
