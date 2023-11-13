package entity

import "gorm.io/gorm"

type Company struct {
	*gorm.Model
	ID           int    `db:"id" gorm:"primary_key"`
	Name         string `db:"name"`
	ContactEmail string `db:"contact_email" gorm:"size:255;"`
	ContactPhone string `db:"contact_phone" gorm:"size:255;"`
}
