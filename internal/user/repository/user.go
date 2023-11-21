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

func (r *Repo) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	res := r.Replica.DB.WithContext(ctx).Where("id = ?", id).First(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get user by id err: %v", res.Error)
	}

	return &user, nil
}

func (r *Repo) ConfirmUser(ctx context.Context, email string) error {
	res := r.Replica.DB.Model(&entity.User{}).WithContext(ctx).Where("email = ?", email).Updates(entity.User{
		IsConfirmed: true,
	})
	if res.Error != nil {
		return fmt.Errorf("failed to confirm user err: %v", res.Error)
	}

	return nil
}
