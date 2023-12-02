package pickup

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type UseCase interface {
	OrderUseCase
	TransactionUseCase
	ProductUseCase
	OrderItemUseCase
}

type OrderUseCase interface {
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	CreateOrder(ctx context.Context, request CreateOrderRequest) (int, error)
	DeleteOrder(ctx context.Context, code string) (string, error)
	GetOrders(ctx context.Context, request GetOrdersQuery) (*[]entity.Order, error)
	PickupOrder(ctx context.Context, code string) error
	ReceiveOrder(ctx context.Context, code string) error
	CancelOrder(ctx context.Context, code string) error
	RefundOrder(ctx context.Context, code string) error
}

type TransactionUseCase interface {
	GetTransactions(ctx context.Context, userID int, query GetTransactionsQuery) (*[]entity.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (int, error)
}

type ProductUseCase interface {
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
}

//type WarehouseUseCase interface {
//	GetWarehouseOrders(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error)
//	GetWarehouse(ctx context.Context, pointID int) (*entity.Warehouse, error)
//	DeleteWarehouseOrder(ctx context.Context, orderID int) error
//	CreateWarehouseOrder(ctx context.Context, order *entity.WarehouseOrder) (int, error)
//}

type OrderItemUseCase interface {
	ReceiveItem(ctx context.Context, orderCode string, productID int) error
	RefundItem(ctx context.Context, orderCode string, request RefundItemRequest) error
}
