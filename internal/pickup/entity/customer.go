package entity

type Customer struct {
	ID        int    `json:"id" db:"id" gorm:"primary_key;"`
	FirstName string `json:"first_name" db:"first_name" gorm:"size:255;"`
	LastName  string `json:"last_name" db:"last_name" gorm:"size:255;"`
	Email     string `json:"email" db:"email" gorm:"size:255;"`
	Phone     string `json:"phone" db:"phone" gorm:"size:15;"`
}
