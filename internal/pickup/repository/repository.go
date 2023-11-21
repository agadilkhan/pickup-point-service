package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Repository interface {
	OrderRepository
	OrderPickupRepository
	CustomerRepository
	CompanyRepository
}

type OrderRepository interface {
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
}

type OrderPickupRepository interface {
	CreateOrderPickup(ctx context.Context, pickup *entity.OrderPickup) (int, error)
}

type CustomerRepository interface {
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
}

type CompanyRepository interface {
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
