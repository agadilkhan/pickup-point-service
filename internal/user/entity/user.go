package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID        int    `db:"id" gorm:"primary_key"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	Login     string `db:"login"`
	Password  string `db:"password"`
}
