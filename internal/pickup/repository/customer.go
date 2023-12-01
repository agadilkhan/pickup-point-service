package repository

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error) {
	var customer entity.Customer

	res := r.replica.DB.WithContext(ctx).Where("id = ?", id).First(&customer)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get customer err: %v", res.Error)
	}

	return &customer, nil
}

func (r *Repo) GetCustomers(ctx context.Context) (*[]entity.Customer, error) {
	var customers []entity.Customer

	res := r.replica.DB.WithContext(ctx).Find(&customers)
	if res.Error != nil {
		return nil, res.Error
	}

	return &customers, nil
}
