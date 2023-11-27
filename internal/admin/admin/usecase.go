package admin

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/admin/entity"
)

type UseCase interface {
	UserUseCase
}

type UserUseCase interface {
	GetUsers(ctx context.Context) (*[]entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	UpdateUser(ctx context.Context, request UpdateUserRequest) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}
