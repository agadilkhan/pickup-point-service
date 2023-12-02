package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	var product entity.Product

	res := r.replica.DB.WithContext(ctx).Where("id = ?", id).First(&product)
	if res.Error != nil {
		return nil, res.Error
	}

	return &product, nil
}
