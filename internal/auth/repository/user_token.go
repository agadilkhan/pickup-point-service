package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

func (r *Repo) CreateUserToken(ctx context.Context, userToken entity.UserToken) error {
	res := r.main.DB.WithContext(ctx).Create(&userToken)
	if res.Error != nil {
		return fmt.Errorf("failed to create user token err: %w", res.Error)
	}

	return nil
}

func (r *Repo) UpdateUserToken(ctx context.Context, userToken entity.UserToken) error {
	res := r.main.DB.Model(&userToken).WithContext(ctx).Updates(entity.UserToken{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	})

	if res.Error != nil {
		return fmt.Errorf("failed to update user token err: %v", res.Error)
	}

	return nil
}

func (r *Repo) GetUserToken(ctx context.Context, refreshToken string) (*entity.UserToken, error) {
	var token entity.UserToken

	res := r.replica.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&token)
	if res.Error != nil {
		return nil, fmt.Errorf("not found err %v", res.Error)
	}

	return &token, nil
}
