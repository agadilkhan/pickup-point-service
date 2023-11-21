package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetOrderByCode(ctx context.Context, code string) (*entity.Order, error) {
	var o entity.Order

	res := r.replica.DB.WithContext(ctx).Where("code = ?", code).First(&o)
	if res.Error != nil {
		return nil, fmt.Errorf("could not find order by code")
	}

	return &o, nil
}

func (r *Repo) UpdateOrder(ctx context.Context, order *entity.Order) error {
	res := r.main.DB.Model(&order).WithContext(ctx).Updates(entity.Order{
		Status: order.Status,
	})

	if res.Error != nil {
		return fmt.Errorf("failed to update order err: %v", res.Error)
	}

	return nil
}

func (r *Repo) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	res := r.main.DB.WithContext(ctx).Create(order)
	if res.Error != nil {
		return 0, fmt.Errorf("could not create order")
	}

	return order.ID, nil
}
