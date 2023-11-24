package entity

type Company struct {
	ID           int    `json:"id" db:"id" gorm:"primary_key;"`
	Name         string `json:"name" db:"name" gorm:"size:255;"`
	ContactEmail string `json:"contact_email" db:"contact_email" gorm:"size:255;"`
	ContactPhone string `json:"contact_phone" db:"contact_phone" gorm:"size:255;"`
}
