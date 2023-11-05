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
