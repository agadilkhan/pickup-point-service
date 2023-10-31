package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

type Repository interface {
	UserTokenRepository
}

type UserTokenRepository interface {
	CreateUserToken(ctx context.Context, userToken entity.UserToken) error
	UpdateUserToken(ctx context.Context, userToken entity.UserToken) error
}
