package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
)

func (r *Repo) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	res := r.Main.DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to create user err: %v", res.Error)
	}

	return user.ID, nil
}

func (r *Repo) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	var user entity.User

	res := r.Replica.DB.WithContext(ctx).Where("login = ?", login).First(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to user find err: %v", res.Error)
	}

	return &user, nil
}
