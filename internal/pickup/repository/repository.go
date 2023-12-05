package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/pkg/pagination"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Repository interface {
	OrderRepository
	TransactionRepository
	ProductRepository
	OrderItemRepository
}

type OrderRepository interface {
	GetOrders(ctx context.Context, sortOptions pagination.SortOptions, filterOptions pagination.FilterOptions) (*[]entity.Order, error)
	GetOrderByCode(ctx context.Context, code string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (int, error)
	DeleteOrder(ctx context.Context, orderCode string) (string, error)
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (int, error)
	GetTransactions(ctx context.Context, userID int, transactionType string) (*[]entity.Transaction, error)
}

type ProductRepository interface {
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
	GetProducts(ctx context.Context, searchOptions pagination.SearchOptions) (*[]entity.Product, error)
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
