package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/metrics"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) UpdateItem(ctx context.Context, item *entity.OrderItem) (*entity.OrderItem, error) {
	ok, fail := metrics.DatabaseQueryTime("GetItems")
	defer fail()

	res := r.main.DB.WithContext(ctx).Where("id = ?", item.ID).Updates(&item)
	if res.Error != nil {
		return nil, res.Error
	}

	ok()

	return item, nil
}
