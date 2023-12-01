package entity

import "time"

type Warehouse struct {
	ID          int         `json:"id" db:"id" gorm:"primary_key"`
	PointID     int         `json:"point_id"`
	NumOfPlaces int         `json:"num_of_places"`
	Point       PickupPoint `json:"point"`
}

type WarehouseOrder struct {
	ID          int       `json:"id" db:"id" gorm:"primary_key"`
	WarehouseID int       `json:"warehouse_id" db:"warehouse_id"`
	OrderID     int       `json:"order_id" db:"order_id"`
	PlaceNum    int       `json:"place_num" db:"place_num"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Warehouse   Warehouse `json:"warehouse"`
	Order       Order     `json:"order"`
}
