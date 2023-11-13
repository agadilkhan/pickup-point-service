package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/auth/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

type Repository interface {
	UserTokenRepository
}

type UserTokenRepository interface {
	CreateUserToken(ctx context.Context, userToken entity.UserToken) error
	UpdateUserToken(ctx context.Context, userToken entity.UserToken) error
	GetUserToken(ctx context.Context, refreshToken string) (*entity.UserToken, error)
}

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
