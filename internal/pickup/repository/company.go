package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetCompanyByID(ctx context.Context, id int) (*entity.Company, error) {
	var company entity.Company

	res := r.replica.WithContext(ctx).Where("id = ?", id).First(&company)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get company by id err: %v", res.Error)
	}

	return &company, nil
}

func (r *Repo) GetAllCompanies(ctx context.Context) (*[]entity.Company, error) {
	var companies []entity.Company

	res := r.replica.DB.WithContext(ctx).Find(&companies)
	if res.Error != nil {
		return nil, res.Error
	}

	return &companies, nil
}