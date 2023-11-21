package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"gorm.io/gorm"
)

func (r *Repo) CreateOrderPickup(ctx context.Context, pickup *entity.OrderPickup) (int, error) {
	err := r.main.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.WithContext(ctx).Create(&pickup)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Model(&pickup.Order).WithContext(ctx).Updates(entity.Order{
			Status: entity.OrderStatusGiven,
		})
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return pickup.ID, nil
}
