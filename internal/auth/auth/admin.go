package auth

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"github.com/agadilkhan/pickup-point-service/internal/auth/transport"
)

func (s *Service) GetUsers(ctx context.Context) (*[]entity.User, error) {
	resp, err := s.userGrpcTransport.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetUsers err: %v", err)
	}

	var users []entity.User
	//nolint:all
	for _, u := range *resp {
		user := entity.User{
			ID:          int(u.Result.Id),
			RoleID:      int(u.Result.RoleId),
			FirstName:   u.Result.FirstName,
			LastName:    u.Result.LastName,
			Email:       u.Result.Email,
			Phone:       u.Result.Phone,
			Login:       u.Result.Login,
			Password:    u.Result.Password,
			IsConfirmed: u.Result.IsConfirmed,
		}

		users = append(users, user)
	}

	return &users, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	resp, err := s.userGrpcTransport.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetUserByID err: %v", err)
	}

	user := entity.User{
		ID:          int(resp.Id),
		RoleID:      int(resp.RoleId),
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		Phone:       resp.Phone,
		Login:       resp.Login,
		Password:    resp.Password,
		IsConfirmed: resp.IsConfirmed,
	}

	return &user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.Password = s.generatePassword(user.Password)

	updateRequest := transport.UpdateUserRequest{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Login:     user.Login,
		Password:  user.Password,
	}

	resp, err := s.userGrpcTransport.UpdateUser(ctx, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to UpdateUser err: %v", err)
	}

	return &entity.User{
		ID:          int(resp.Id),
		RoleID:      int(resp.RoleId),
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		Phone:       resp.Phone,
		Login:       resp.Login,
		Password:    resp.Password,
		IsConfirmed: resp.IsConfirmed,
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) (int, error) {
	resp, err := s.userGrpcTransport.DeleteUser(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("failed to DeleteUser err: %v", err)
	}

	return int(resp.Id), nil
}
