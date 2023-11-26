package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Repository interface {
	OrderRepository
	PickupRepository
	CustomerRepository
	CompanyRepository
	ProductRepository
	WarehouseRepository
}

type OrderRepository interface {
	GetOrders(ctx context.Context, sort, direction string) (*[]entity.Order, error)
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
}

type PickupRepository interface {
	CreatePickupOrder(ctx context.Context, pickup *entity.PickupOrder) (int, error)
	GetPickupOrders(ctx context.Context, userID int) (*[]entity.PickupOrder, error)
	GetPickupOrderByID(ctx context.Context, userID, pickupOrderID int) (*entity.PickupOrder, error)
	GetAllPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error)
	GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error)
}

type CustomerRepository interface {
	GetAllCustomers(ctx context.Context) (*[]entity.Customer, error)
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
}

type CompanyRepository interface {
	GetAllCompanies(ctx context.Context) (*[]entity.Company, error)
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
}

type ProductRepository interface {
	GetProductByID(ctx context.Context, id int) (*entity.Product, error)
	GetAllProducts(ctx context.Context) (*[]entity.Product, error)
}

type WarehouseRepository interface {
	GetAllWarehouses(ctx context.Context) (*[]entity.Warehouse, error)
	GetWarehouseByID(ctx context.Context, id int) (*entity.Warehouse, error)
	CreateWarehouseOrder(ctx context.Context, warehouseOrder *entity.OrderWarehouse) error
	GetWarehouseOrders(ctx context.Context, warehouseID int) (*[]entity.OrderWarehouse, error)
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
