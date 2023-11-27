package grpc

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
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
			RoleId:    int64(user.RoleID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Login:     user.Login,
			Password:  user.Password,
		},
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := entity.User{
		RoleID:    2,
		FirstName: request.Request.FirstName,
		LastName:  request.Request.LastName,
		Email:     request.Request.Email,
		Phone:     request.Request.Phone,
		Login:     request.Request.Login,
		Password:  request.Request.Password,
	}

	id, err := s.repo.CreateUser(ctx, &user)
	if err != nil {
		s.logger.Errorf("failed to CreateUser err: %v", err)
		return nil, fmt.Errorf("CreateUser err: %v", err)
	}

	s.logger.Info("CreateUser success")

	return &pb.CreateUserResponse{
		Id: int64(id),
	}, nil
}

func (s *Service) GetUserByID(ctx context.Context, request *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := s.repo.GetUserByID(ctx, int(request.Id))
	if err != nil {
		s.logger.Errorf("failed to GetUserByID err: %v", err)
		return nil, fmt.Errorf("GetUserByID err: %v", err)
	}

	return &pb.GetUserByIDResponse{
		Result: &pb.User{
			Id:          int64(user.ID),
			RoleId:      int64(user.RoleID),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			Phone:       user.Phone,
			Login:       user.Login,
			Password:    user.Password,
			IsConfirmed: user.IsConfirmed,
		},
	}, nil
}

func (s *Service) ConfirmUser(ctx context.Context, request *pb.ConfirmUserRequest) (*pb.ConfirmUserResponse, error) {
	err := s.repo.ConfirmUser(ctx, request.Email)
	if err != nil {
		s.logger.Errorf("failed to ConfirmUser err: %v", err)
		return nil, fmt.Errorf("ConfirmUser err: %v", err)
	}

	return &pb.ConfirmUserResponse{}, nil
}

func (s *Service) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	updateRequest := entity.User{
		ID:          int(request.Request.Id),
		RoleID:      2,
		FirstName:   request.Request.FirstName,
		LastName:    request.Request.LastName,
		Email:       request.Request.Email,
		Phone:       request.Request.Phone,
		Login:       request.Request.Login,
		Password:    request.Request.Password,
		IsConfirmed: false,
		IsDeleted:   false,
	}

	user, err := s.repo.UpdateUser(ctx, &updateRequest)
	if err != nil {
		s.logger.Errorf("failed to UpdateUser: %v", err)
		return nil, fmt.Errorf("UpdateUser err: %v", err)
	}

	return &pb.UpdateUserResponse{
		Result: &pb.User{
			Id:          int64(user.ID),
			RoleId:      int64(user.RoleID),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			Phone:       user.Phone,
			Login:       user.Login,
			Password:    user.Password,
			IsConfirmed: user.IsConfirmed,
		},
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id, err := s.repo.DeleteUser(ctx, int(request.Id))
	if err != nil {
		s.logger.Errorf("failed to DeleteUser err: %v", err)
		return nil, fmt.Errorf("DeleteUser err: %v", err)
	}

	return &pb.DeleteUserResponse{
		Id: int64(id),
	}, nil
}

//func (s *Service) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
//	users, err := s.repo.GetUsers(ctx)
//	if err != nil {
//		s.logger.Errorf("failed to GetUsers: %v", err)
//		return nil, fmt.Errorf("failed to GetUsers err: %v", err)
//	}
//
//	for _, user := range *users {
//
//	}
//
//}
