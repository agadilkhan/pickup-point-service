package dto

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"time"
)

type OrderDTO struct {
	ID          int         `json:"id" gorm:"primary_key;"`
	CustomerID  int         `json:"customer_id"`
	Customer    entity.Customer    `json:"customer" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	CompanyID   int         `json:"company_id"`
	Company     entity.Company     `json:"company" gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
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
