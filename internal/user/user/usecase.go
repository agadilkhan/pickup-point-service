package user

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
)

type UseCase interface {
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
}
