package pickup

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

type Order interface {
	GetOrderByID(ctx context.Context, id int) (*entity.Order, error)
}
