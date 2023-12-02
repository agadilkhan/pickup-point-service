package entity

type Customer struct {
	ID        int    `json:"id" gorm:"primary_key;"`
	FirstName string `json:"first_name" gorm:"size:255;"`
	LastName  string `json:"last_name" gorm:"size:255;"`
	Email     string `json:"email" gorm:"size:255;"`
	Phone     string `json:"phone" gorm:"size:15;"`
}
