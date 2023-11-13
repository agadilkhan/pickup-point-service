package entity

import "gorm.io/gorm"

type Product struct {
	*gorm.Model
	ID          int     `db:"id" gorm:"primary_key"`
	Name        string  `db:"name" gorm:"size:255;"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
}
