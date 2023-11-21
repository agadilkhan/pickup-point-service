package entity

type Role struct {
	ID          int    `db:"id" gorm:"primary_key;"`
	Name        string `db:"name" gorm:"size:50; not null;"`
	Description string `db:"description"`
}
