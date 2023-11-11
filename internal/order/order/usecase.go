package order

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/order/entity"
)

type Order interface {
	GetOrderByID(ctx context.Context, id int) (*entity.Order, error)
}
