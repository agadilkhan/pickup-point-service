package entity

type Product struct {
	ID          int     `json:"id" gorm:"primary_key;"`
	Name        string  `json:"name" gorm:"size:255;"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
