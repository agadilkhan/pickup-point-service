package entity

type Product struct {
	ID          int     `json:"id" db:"id" gorm:"primary_key;"`
	Name        string  `json:"name" db:"name" gorm:"size:255;"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
}
