package entity

import "time"

type OrderStatus string
type PaymentStatus string

var (
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusReady      OrderStatus = "ready to pickup"
	OrderStatusGiven      OrderStatus = "given"
)

var (
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusNotPaid PaymentStatus = "not paid"
)

type Order struct {
	ID            int           `db:"id" gorm:"primary_key;"`
	CustomerID    int           `db:"customer_id"`
	CompanyID     int           `db:"company_id"`
	PointID       int           `db:"point_id"`
	Code          string        `db:"code" gorm:"size:50;"`
	Status        OrderStatus   `db:"status" gorm:"size:50;"`
	PaymentStatus PaymentStatus `db:"payment_status" gorm:"size:50;"`
	TotalAmount   float64       `db:"total_amount"`
	CreatedAt     time.Time     `db:"created_at" gorm:"default:now();"`
	UpdatedAt     time.Time     `db:"updated_at" gorm:"default:now();"`
	OrderItems    []OrderItem
	Customer      Customer    `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Company       Company     `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Point         PickupPoint `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type OrderItem struct {
	ID        int     `db:"id" gorm:"primary_key;"`
	OrderID   int     `db:"order_id"`
	ProductID int     `db:"product_id"`
	Quantity  int     `db:"quantity"`
	SubTotal  float64 `db:"sub_total"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
