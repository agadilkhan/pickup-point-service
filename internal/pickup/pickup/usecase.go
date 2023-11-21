package pickup

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type UseCase interface {
	Pickup(ctx context.Context, code string) error
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
}
