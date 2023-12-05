package pickup

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

// swagger:model CreateOrderRequest
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

//swagger:model RefundItemRequest
type RefundItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type GetTransactionsQuery struct {
	TransactionType string
}
