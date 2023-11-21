package entity

type Product struct {
	ID          int     `db:"id" gorm:"primary_key;"`
	Code        string  `db:"code"`
	Name        string  `db:"name" gorm:"size:255;"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
}
