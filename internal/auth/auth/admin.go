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

func (s *Service) UpdateUser(ctx context.Context, request UpdateUserRequest) (*entity.User, error) {
	request.Password = s.generatePassword(request.Password)

	updateRequest := transport.UpdateUserRequest{
		ID:        request.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Login:     request.Login,
		Password:  request.Password,
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
