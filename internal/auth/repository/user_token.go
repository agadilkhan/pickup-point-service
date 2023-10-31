package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

type Repo struct {
	main    *postgres.Db
	replica *postgres.Db
}

func NewRepository(main *postgres.Db, replica *postgres.Db) *Repo {
	return &Repo{
		main,
		replica,
	}
}

func (r *Repo) CreateUserToken(ctx context.Context, userToken entity.UserToken) error {
	res := r.main.DB.WithContext(ctx).Create(&userToken)
	if res.Error != nil {
		return fmt.Errorf("failed to create user token err: %w", res.Error)
	}

	return nil
}

func (r *Repo) UpdateUserToken(ctx context.Context, userToken entity.UserToken) error {
	return nil
}
