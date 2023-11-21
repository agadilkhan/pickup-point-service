package entity

import "time"

type OrderPickup struct {
	ID        int       `db:"id" gorm:"primary_key;"`
	UserID    int       `db:"user_id"`
	OrderID   int       `db:"order_id"`
	CreatedAt time.Time `db:"created_at" gorm:"default:now();"`
	UpdatedAt time.Time `db:"updated_at" gorm:"default:now();"`
	Order     Order     `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type PickupPoint struct {
	ID      int    `db:"id" gorm:"primary_key;"`
	Name    string `db:"name" gorm:"size:255;"`
	Address string `db:"address" gorm:"size:255"`
}
