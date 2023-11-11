package entity

import "gorm.io/gorm"

type Customer struct {
	*gorm.Model
	ID int `db:"id" gorm:"primary_key"`
}
