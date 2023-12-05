package auth

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

type UseCase interface {
	TokenUseCase
	UserUseCase
	AdminUseCase
}

type TokenUseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error)
	RenewToken(ctx context.Context, refreshToken string) (*JWTUserToken, error)
}

type UserUseCase interface {
	Register(ctx context.Context, request CreateUserRequest) (int, error)
	ConfirmUser(cxt context.Context, code ConfirmUserRequest) error
}

type AdminUseCase interface {
	GetUsers(ctx context.Context) (*[]entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}
