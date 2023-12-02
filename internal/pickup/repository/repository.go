package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Repository interface {
	OrderRepository
	TransactionRepository
	//WarehouseRepository
	ProductRepository
	OrderItemRepository
}

type OrderRepository interface {
	GetOrders(ctx context.Context, sort, direction string) (*[]entity.Order, error)
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
	DeleteOrder(ctx context.Context, orderCode string) (string, error)
}

//type WarehouseRepository interface {
//	GetWarehouse(ctx context.Context, pointID int) (*entity.Warehouse, error)
//	CreateWarehouseOrder(ctx context.Context, warehouseOrder *entity.WarehouseOrder) (int, error)
//	GetWarehouseOrders(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error)
//	DeleteWarehouseOrder(ctx context.Context, orderID int) error
//}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (int, error)
	GetTransactions(ctx context.Context, userID int, transactionType string) (*[]entity.Transaction, error)
}

type ProductRepository interface {
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
}

type OrderItemRepository interface {
	UpdateItem(ctx context.Context, item *entity.OrderItem) (*entity.OrderItem, error)
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
