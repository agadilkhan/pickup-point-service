package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/user/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
)

type Repository interface {
	UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (int, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	ConfirmUser(ctx context.Context, email string) error
	GetUsers(ctx context.Context) (*[]entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
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
