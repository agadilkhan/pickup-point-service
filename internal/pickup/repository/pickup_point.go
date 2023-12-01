package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error) {
	var points []entity.PickupPoint

	res := r.replica.DB.WithContext(ctx).Find(&points)
	if res.Error != nil {
		return nil, res.Error
	}

	return &points, nil
}

func (r *Repo) GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error) {
	var point entity.PickupPoint

	res := r.replica.DB.WithContext(ctx).Where("id = ?", id).First(&point)
	if res.Error != nil {
		return nil, res.Error
	}

	return &point, nil
}
