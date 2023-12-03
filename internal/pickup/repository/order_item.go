package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) UpdateItem(ctx context.Context, item *entity.OrderItem) (*entity.OrderItem, error) {
	res := r.main.DB.WithContext(ctx).Where("id = ?", item.ID).Updates(&item)
	if res.Error != nil {
		return nil, res.Error
	}

	return item, nil
}
