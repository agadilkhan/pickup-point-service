package pickup

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type UseCase interface {
	PickupPointUseCase
	OrderUseCase
	CustomerUseCase
	CompanyUseCase
	WarehouseUseCase
	ProductUseCase
	TransactionUseCase
}

type OrderUseCase interface {
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	CreateOrder(ctx context.Context, request CreateOrderRequest) (int, error)
	GetOrders(ctx context.Context, request GetOrdersQuery) (*[]entity.Order, error)
	PickupOrder(ctx context.Context, code string) error
	ReceiveOrder(ctx context.Context, request ReceiveOrderRequest) (int, error)
}

type PickupPointUseCase interface {
	GetPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error)
	GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error)
}

type CustomerUseCase interface {
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
	GetCustomers(ctx context.Context) (*[]entity.Customer, error)
}

type CompanyUseCase interface {
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
	GetCompanies(ctx context.Context, query GetCompaniesQuery) (*[]entity.Company, error)
}

type ProductUseCase interface {
	GetProducts(ctx context.Context) (*[]entity.Product, error)
	GetProductByID(ctx context.Context, id int) (*entity.Product, error)
}

type WarehouseUseCase interface {
	GetWarehouses(ctx context.Context) (*[]entity.Warehouse, error)
	GetWarehouseByID(ctx context.Context, id int) (*entity.Warehouse, error)
}

type TransactionUseCase interface {
	GetTransactions(ctx context.Context, userID int, query GetTransactionsQuery) (*[]entity.Transaction, error)
}
