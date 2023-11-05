package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
)

func (r *Repo) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	return 0, nil
}

func (r *Repo) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	var user *entity.User

	res := r.Replica.DB.WithContext(ctx).Where("login = ?", login).Find(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to user find err: %v", res.Error)
	}

	return user, nil
}
