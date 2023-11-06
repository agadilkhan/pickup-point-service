package user

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
	"github.com/agadilkhan/pickup-point-service/internal/user/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) UseCase {
	return &Service{repo: repo}
}

func (s *Service) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	user, err := s.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("GetUserByLogin err: %v", err)
	}

	return user, nil
}

func (s *Service) CreateUser(ctx context.Context, request CreateUserRequest) (int, error) {
	user := entity.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Login:     request.Login,
		Password:  request.Password,
	}

	userID, err := s.repo.CreateUser(ctx, &user)
	if err != nil {
		return 0, fmt.Errorf("CreateUser request err: %v", err)
	}

	return userID, nil
}
