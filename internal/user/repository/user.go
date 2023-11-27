package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
)

func (r *Repo) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	res := r.main.DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}

func (r *Repo) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	var user entity.User

	res := r.replica.DB.WithContext(ctx).Where("login = ? AND is_deleted = false", login).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (r *Repo) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	res := r.replica.DB.WithContext(ctx).Where("id = ? AND is_deleted = false", id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (r *Repo) ConfirmUser(ctx context.Context, email string) error {
	res := r.replica.DB.Model(&entity.User{}).WithContext(ctx).Where("email = ?", email).Updates(entity.User{
		IsConfirmed: true,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repo) GetUsers(ctx context.Context) (*[]entity.User, error) {
	var users []entity.User

	res := r.replica.DB.WithContext(ctx).Where("is_deleted = false").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return &users, nil
}

func (r *Repo) UpdateUser(ctx context.Context, updatedUser *entity.User) (*entity.User, error) {
	res := r.main.DB.WithContext(ctx).Model(&updatedUser).Updates(&updatedUser)
	if res.Error != nil {
		return nil, res.Error
	}

	return updatedUser, nil
}

func (r *Repo) DeleteUser(ctx context.Context, id int) (int, error) {
	var user entity.User

	res := r.replica.DB.WithContext(ctx).Where("id = ? AND is_deleted = false", id).First(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	res = r.main.WithContext(ctx).Model(&user).Updates(entity.User{
		IsDeleted: true,
	})
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}
