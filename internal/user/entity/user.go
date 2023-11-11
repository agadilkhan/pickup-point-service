package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID        int    `db:"id" gorm:"primary_key"`
	RoleID    int    `db:"role_id" gorm:"not null"`
	FirstName string `db:"first_name" gorm:"size:255;"`
	LastName  string `db:"last_name" gorm:"size:255;"`
	Email     string `db:"email" gorm:"size:255;"`
	Phone     string `db:"phone" gorm:"size:15;"`
	Login     string `db:"login" gorm:"size:255"`
	Password  string `db:"password" gorm:"size:255;"`
	Role      Role   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
