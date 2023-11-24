package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Repository interface {
	OrderRepository
	PickupOrderRepository
	CustomerRepository
	CompanyRepository
}

type OrderRepository interface {
	GetOrders(ctx context.Context, sort, direction string) (*[]entity.Order, error)
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
}

type PickupOrderRepository interface {
	CreatePickupOrder(ctx context.Context, pickup *entity.PickupOrder) (int, error)
	GetPickupOrders(ctx context.Context, userID int) (*[]entity.PickupOrder, error)
	GetPickupOrderByID(ctx context.Context, userID, pickupOrderID int) (*entity.PickupOrder, error)
}

type CustomerRepository interface {
	GetAllCustomers(ctx context.Context) (*[]entity.Customer, error)
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
}

type CompanyRepository interface {
	GetAllCompanies(ctx context.Context) (*[]entity.Company, error)
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
}

type Repo struct {
	main    *postgres.Db
	replica *postgres.Db
}

func NewRepository(main, replica *postgres.Db) *Repo {
	return &Repo{
		main:    main,
		replica: replica,
	}
}
