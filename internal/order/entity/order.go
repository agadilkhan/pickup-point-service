package entity

import "gorm.io/gorm"

type Order struct {
	*gorm.Model
	ID int `db:"id" gorm:"primary_key"`
}
