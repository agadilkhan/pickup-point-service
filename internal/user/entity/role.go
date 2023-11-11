package entity

import "gorm.io/gorm"

type Role struct {
	*gorm.Model
	ID          int    `db:"id" gorm:"primary_key"`
	Name        string `db:"name" gorm:"size:50; not null;"`
	Description string `db:"description"`
}
