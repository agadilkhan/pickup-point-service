package entity

import (
	"gorm.io/gorm"
)

type UserToken struct {
	*gorm.Model
	ID           int    `db:"id" gorm:"primary_key"`
	Token        string `db:"token"`
	RefreshToken string `db:"refresh_token"`
	UserID       int    `db:"user_id"`
}
