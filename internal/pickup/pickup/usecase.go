package pickup

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type UseCase interface {
	PickupOrderUseCase
	OrderUseCase
	CustomerUseCase
	CompanyUseCase
}

type OrderUseCase interface {
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
	GetOrders(ctx context.Context, request GetAllOrdersRequest) (*[]entity.Order, error)
}

type PickupOrderUseCase interface {
	Pickup(ctx context.Context, code string) error
	GetPickupOrders(ctx context.Context, userID int) (*[]entity.PickupOrder, error)
	GetPickupOrderByID(ctx context.Context, request GetPickupOrderByIDRequest) (*entity.PickupOrder, error)
}

type CustomerUseCase interface {
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
}

type CompanyUseCase interface {
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
}
