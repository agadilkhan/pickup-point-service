package pickup

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/pkg/pagination"
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
	GetOrders(ctx context.Context, sortOptions pagination.SortOptions, filterOptions pagination.FilterOptions) (*[]entity.Order, error)
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
	GetProducts(ctx context.Context, searchOptions pagination.SearchOptions) (*[]entity.Product, error)
}

type OrderItemUseCase interface {
	ReceiveItem(ctx context.Context, orderCode string, productID int) error
	RefundItem(ctx context.Context, orderCode string, request RefundItemRequest) error
}
