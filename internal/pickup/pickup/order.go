package pickup

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) GetOrderByID(ctx context.Context, id int) (*entity.Order, error) {
	return nil, nil
}
