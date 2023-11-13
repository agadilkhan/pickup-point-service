package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Order interface {
	GetOrderByID(ctx context.Context, id int) (*entity.Order, error)
}

type Customer interface {
}

type Repository struct {
	Order
	Customer
}

func NewRepository(main *postgres.Db, replica *postgres.Db) *Repository {
	return &Repository{
		Order: NewOrderRepo(main, replica),
	}
}
