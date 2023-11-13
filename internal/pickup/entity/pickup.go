package entity

import "gorm.io/gorm"

type OrderPickup struct {
	*gorm.Model
	ID      int         `db:"id" gorm:"primary_key"`
	UserID  int         `db:"user_id"`
	OrderID int         `db:"order_id"`
	PointID int         `db:"point_id"`
	Point   PickupPoint `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type PickupPoint struct {
	*gorm.Model
	ID      int    `db:"id" gorm:"primary_key"`
	Name    string `db:"name" gorm:"size:255;"`
	Address string `db:"address" gorm:"size:255"`
}
