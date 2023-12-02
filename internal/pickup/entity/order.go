package entity

import "time"

type OrderStatus string

const (
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefund     OrderStatus = "refund"
	OrderStatusReady      OrderStatus = "ready to pickup"
	OrderStatusGiven      OrderStatus = "given"
)

type Order struct {
	ID          int         `json:"id" gorm:"primary_key;"`
	CustomerID  int         `json:"customer_id"`
	Customer    Customer    `json:"customer" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	CompanyID   int         `json:"company_id"`
	Company     Company     `json:"company" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	PointID     int         `json:"point_id"`
	Point       PickupPoint `json:"point" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	OrderItems  []OrderItem `json:"items"`
	Code        string      `json:"code" gorm:"size:50;"`
	Status      OrderStatus `json:"status" gorm:"size:50;"`
	IsPaid      bool        `json:"is_paid"`
	TotalAmount float64     `json:"total_amount"`
	CreatedAt   time.Time   `json:"created_at" gorm:"default:now();"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"default:now();"`
}

type OrderItem struct {
	ID          int     `json:"id" gorm:"primary_key;"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	Product     Product `json:"product" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	Quantity    int     `json:"quantity"`
	SubTotal    float64 `json:"sub_total"`
	IsAccept    bool    `json:"is_accept"`
	NumOfRefund int     `json:"num_of_refund"`
}
