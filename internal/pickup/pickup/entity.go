package pickup

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"time"
)

type GetOrdersQuery struct {
	Sort      string
	Direction string
}

type GetPickupOrderByIDRequest struct {
	UserID        int
	PickupOrderID int
}

type CreateOrderRequest struct {
	CustomerID    int                  `json:"customer_id"`
	CompanyID     int                  `json:"company_id"`
	PointID       int                  `json:"point_id"`
	Status        entity.OrderStatus   `json:"status"`
	PaymentStatus entity.PaymentStatus `json:"payment_status"`
	Items         []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"items"`
}

type ReceiveOrderRequest struct {
	WarehouseID int    `json:"warehouse_id"`
	OrderCode   string `json:"order_code"`
	Items       []struct {
		ProductID int `json:"product_id"`
	} `json:"items"`
}

type GetCompaniesQuery struct {
	Name string
}

type GetPickupOrdersQuery struct {
	StartDate time.Time
	EndDate   time.Time
}

type GetTransactionsQuery struct {
	TransactionType string
}
