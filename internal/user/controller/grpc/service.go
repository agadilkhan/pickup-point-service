package grpc

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/user/repository"
	pb "github.com/agadilkhan/pickup-point-service/pkg/protobuf/userservice/gw"
	"go.uber.org/zap"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	logger *zap.SugaredLogger
	repo   repository.Repository
}

func NewService(logger *zap.SugaredLogger, repo repository.Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) GetUserByLogin(ctx context.Context, request *pb.GetUserByLoginRequest) (*pb.GetUserByLoginResponse, error) {
	user, err := s.repo.GetUserByLogin(ctx, request.Login)
	if err != nil {
		s.logger.Errorf("failed to GetUserByLogin err: %v", err)
		return nil, fmt.Errorf("GetUserByLogin err: %v", err)
	}

	return &pb.GetUserByLoginResponse{
		Result: &pb.User{
			Id:        int64(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Login:     user.Login,
			Password:  user.Password,
		},
	}, nil
}
