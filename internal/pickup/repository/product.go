package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetProductByID(ctx context.Context, id int) (*entity.Product, error) {
	var product entity.Product

	res := r.replica.DB.WithContext(ctx).Where("id = ?", id).First(&product)
	if res.Error != nil {
		return nil, res.Error
	}

	return &product, nil
}

func (r *Repo) GetProducts(ctx context.Context) (*[]entity.Product, error) {
	var products []entity.Product

	res := r.replica.DB.WithContext(ctx).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}

	return &products, nil
}
