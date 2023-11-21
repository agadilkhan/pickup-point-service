package entity

type Company struct {
	ID           int    `db:"id" gorm:"primary_key;"`
	Name         string `db:"name" gorm:"size:255;"`
	ContactEmail string `db:"contact_email" gorm:"size:255;"`
	ContactPhone string `db:"contact_phone" gorm:"size:255;"`
}
