package entity

import "gorm.io/gorm"

type OrderStatus string

var (
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	*gorm.Model
	ID            int         `db:"id" gorm:"primary_key"`
	CustomerID    int         `db:"customer_id"`
	CompanyID     int         `db:"company_id"`
	Code          int         `db:"code"`
	Status        OrderStatus `db:"status"`
	PaymentStatus bool        `db:"payment_status"`
	TotalAmount   float64     `db:"total_amount"`
	Pickup        OrderPickup
	OrderItems    []OrderItem
	Customer      Customer `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Company       Company  `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type OrderItem struct {
	*gorm.Model
	ID        int     `db:"id" gorm:"primary_key"`
	OrderID   int     `db:"order_id"`
	ProductID int     `db:"product_id"`
	Quantity  int     `db:"quantity"`
	SubTotal  float64 `db:"sub_total"`
	Order     Order   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
