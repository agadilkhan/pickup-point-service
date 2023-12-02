package entity

type Company struct {
	ID           int    `json:"id" gorm:"primary_key;"`
	Name         string `json:"name" gorm:"size:255;"`
	ContactEmail string `json:"contact_email" gorm:"size:255;"`
	ContactPhone string `json:"contact_phone" gorm:"size:255;"`
}
