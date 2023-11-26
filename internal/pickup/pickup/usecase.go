package pickup

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type UseCase interface {
	PickupUseCase
	OrderUseCase
	CustomerUseCase
	CompanyUseCase
	WarehouseUseCase
	ProductUseCase
}

type OrderUseCase interface {
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	CreateOrder(ctx context.Context, request CreateOrderRequest) (int, error)
	GetOrders(ctx context.Context, request GetAllOrdersRequest) (*[]entity.Order, error)
	ReceiveOrder(ctx context.Context, request ReceiveOrderRequest) (int, error)
}

type PickupUseCase interface {
	Pickup(ctx context.Context, code string) error
	GetPickupOrders(ctx context.Context, userID int) (*[]entity.PickupOrder, error)
	GetPickupOrderByID(ctx context.Context, request GetPickupOrderByIDRequest) (*entity.PickupOrder, error)
	GetAllPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error)
	GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error)
}

type CustomerUseCase interface {
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
	GetAllCustomers(ctx context.Context) (*[]entity.Customer, error)
}

type CompanyUseCase interface {
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
	GetAllCompanies(ctx context.Context) (*[]entity.Company, error)
}

type ProductUseCase interface {
	GetAllProducts(ctx context.Context) (*[]entity.Product, error)
	GetProductByID(ctx context.Context, id int) (*entity.Product, error)
}

type WarehouseUseCase interface {
	GetAllWarehouses(ctx context.Context) (*[]entity.Warehouse, error)
	GetWarehouseByID(ctx context.Context, id int) (*entity.Warehouse, error)
}
