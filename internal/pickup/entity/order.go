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
	ID            int           `json:"id" db:"id" gorm:"primary_key;"`
	CustomerID    int           `json:"customer_id" db:"customer_id"`
	CompanyID     int           `json:"company_id" db:"company_id"`
	PointID       int           `json:"point_id" db:"point_id"`
	Code          string        `json:"code" db:"code" gorm:"size:50;"`
	Status        OrderStatus   `json:"status" db:"status" gorm:"size:50;"`
	PaymentStatus PaymentStatus `json:"payment_status" db:"payment_status" gorm:"size:50;"`
	TotalAmount   float64       `json:"total_amount" db:"total_amount"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at" gorm:"default:now();"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at" gorm:"default:now();"`
	OrderItems    []OrderItem   `json:"items"`
	Customer      Customer      `json:"customer" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Company       Company       `json:"company" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Point         PickupPoint   `json:"point" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type OrderItem struct {
	ID        int     `json:"id" db:"id" gorm:"primary_key;"`
	OrderID   int     `json:"order_id" db:"order_id"`
	ProductID int     `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	SubTotal  float64 `json:"sub_total" db:"sub_total"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
