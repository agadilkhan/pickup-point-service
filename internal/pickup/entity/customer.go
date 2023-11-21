package entity

type Customer struct {
	ID        int    `db:"id" gorm:"primary_key;"`
	FirstName string `db:"first_name" gorm:"size:255;"`
	LastName  string `db:"last_name" gorm:"size:255;"`
	Email     string `db:"email" gorm:"size:255;"`
	Phone     string `db:"phone" gorm:"size:15;"`
}
