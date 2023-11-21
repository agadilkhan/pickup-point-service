package entity

import "time"

type User struct {
	ID          int       `db:"id" gorm:"primary_key;"`
	RoleID      int       `db:"role_id" gorm:"not null;"`
	FirstName   string    `db:"first_name" gorm:"size:255;"`
	LastName    string    `db:"last_name" gorm:"size:255;"`
	Email       string    `db:"email" gorm:"size:255;"`
	Phone       string    `db:"phone" gorm:"size:15;"`
	Login       string    `db:"login" gorm:"size:255"`
	Password    string    `db:"password" gorm:"size:255;"`
	IsConfirmed bool      `db:"is_confirmed"`
	CreatedAt   time.Time `db:"created_at" gorm:"default:now();"`
	UpdatedAt   time.Time `db:"updated_at" gorm:"default:now();"`
	Role        Role      `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
