package entity

import "time"

type UserToken struct {
	ID           int       `db:"id"`
	Token        string    `db:"token"`
	RefreshToken string    `db:"refresh_token"`
	UserID       int       `db:"user_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
