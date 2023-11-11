package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/order/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/order/entity"
)

type Order interface {
	GetOrderByID(ctx context.Context, id int) (*entity.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(main *postgres.Db, replica *postgres.Db) *Repository {
	return &Repository{
		Order: NewOrderRepo(main, replica),
	}
}
