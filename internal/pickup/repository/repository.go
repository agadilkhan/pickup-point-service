package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Repository interface {
	OrderRepository
	PickupPointRepository
	CustomerRepository
	CompanyRepository
	ProductRepository
	WarehouseRepository
	Transaction
}

type OrderRepository interface {
	GetOrders(ctx context.Context, sort, direction string) (*[]entity.Order, error)
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
}

type PickupPointRepository interface {
	GetPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error)
	GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error)
}

type CustomerRepository interface {
	GetCustomers(ctx context.Context) (*[]entity.Customer, error)
	GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error)
}

type CompanyRepository interface {
	GetCompanies(ctx context.Context, name string) (*[]entity.Company, error)
	GetCompanyByID(ctx context.Context, id int) (*entity.Company, error)
}

type ProductRepository interface {
	GetProductByID(ctx context.Context, id int) (*entity.Product, error)
	GetProducts(ctx context.Context) (*[]entity.Product, error)
}

type WarehouseRepository interface {
	GetWarehouses(ctx context.Context) (*[]entity.Warehouse, error)
	GetWarehouseByID(ctx context.Context, id int) (*entity.Warehouse, error)
	CreateWarehouseOrder(ctx context.Context, warehouseOrder *entity.WarehouseOrder) (int, error)
	GetWarehouseOrdersByWarehouseID(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error)
	DeleteWarehouseOrderByOrderID(ctx context.Context, orderID int) error
}

type Transaction interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (int, error)
	GetTransactions(ctx context.Context, userID int, transactionType string) (*[]entity.Transaction, error)
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
