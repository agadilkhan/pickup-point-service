package entity

type PickupPoint struct {
	ID      int    `json:"id" db:"id" gorm:"primary_key;"`
	Name    string `json:"name" db:"name" gorm:"size:255;"`
	Address string `json:"address" db:"address" gorm:"size:255"`
}
