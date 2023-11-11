package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/order/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/order/entity"
)

type OrderRepo struct {
	Main    *postgres.Db
	Replica *postgres.Db
}

func NewOrderRepo(main *postgres.Db, replica *postgres.Db) *OrderRepo {
	return &OrderRepo{
		main,
		replica,
	}
}

func (o *OrderRepo) GetOrderByID(ctx context.Context, id int) (*entity.Order, error) {
	return nil, nil
}
