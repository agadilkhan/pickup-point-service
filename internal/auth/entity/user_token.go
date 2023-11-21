package entity

import (
	"time"
)

type UserToken struct {
	ID           int       `db:"id" gorm:"primary_key;"`
	UserID       int       `db:"user_id"`
	Token        string    `db:"token"`
	RefreshToken string    `db:"refresh_token"`
	CreatedAt    time.Time `db:"created_at" gorm:"default:now();"`
	UpdatedAt    time.Time `db:"updated_at" gorm:"default:now();"`
}
