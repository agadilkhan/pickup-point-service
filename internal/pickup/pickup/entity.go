package pickup

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type GetOrdersQuery struct {
	Sort      string
	Direction string
}

type CreateOrderRequest struct {
	CustomerID int                `json:"customer_id"`
	CompanyID  int                `json:"company_id"`
	PointID    int                `json:"point_id"`
	Status     entity.OrderStatus `json:"status"`
	IsPaid     bool               `json:"is_paid"`
	Items      []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"items"`
}

type RefundItemRequest struct {
	ProductID int
	Quantity  int
}

type GetTransactionsQuery struct {
	TransactionType string
}
