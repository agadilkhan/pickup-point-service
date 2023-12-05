package grpc

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/user/memory"

	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
	"github.com/agadilkhan/pickup-point-service/internal/user/repository"
	pb "github.com/agadilkhan/pickup-point-service/pkg/protobuf/userservice/gw"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	logger     *zap.SugaredLogger
	repo       repository.Repository
	userMemory *memory.UserMemory
}

func NewService(logger *zap.SugaredLogger, repo repository.Repository, userMemory *memory.UserMemory) *Service {
	return &Service{
		logger:     logger,
		repo:       repo,
		userMemory: userMemory,
	}
}

func (s *Service) GetUserByLogin(ctx context.Context, request *pb.GetUserByLoginRequest) (*pb.GetUserByLoginResponse, error) {
	var user *entity.User

	memoryUser := s.userMemory.GetUserByLogin(request.Login)
	if memoryUser != nil {
		user = memoryUser
	} else {
		repoUser, err := s.repo.GetUserByLogin(ctx, request.Login)
		if err != nil {
			s.logger.Errorf("GetUserByLogin err: %v", err)
			return nil, fmt.Errorf("failed to GetUserByLogin err: %v", err)
		}

		user = repoUser
	}

	s.logger.Infof("GetUserByLogin success")

	return &pb.GetUserByLoginResponse{
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

func (s *Service) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := entity.User{
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

func (s *Service) ConfirmUser(ctx context.Context, request *pb.ConfirmUserRequest) (*emptypb.Empty, error) {
	err := s.repo.ConfirmUser(ctx, request.Email)
	if err != nil {
		s.logger.Errorf("failed to ConfirmUser err: %v", err)
		return nil, fmt.Errorf("ConfirmUser err: %v", err)
	}

	s.logger.Infof("ConfrimUser success")

	return &emptypb.Empty{}, nil
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
		IsConfirmed: true,
		IsDeleted:   false,
	}

	user, err := s.repo.UpdateUser(ctx, &updateRequest)
	if err != nil {
		s.logger.Errorf("failed to UpdateUser: %v", err)
		return nil, fmt.Errorf("UpdateUser err: %v", err)
	}

	s.logger.Infof("UpdateUser success")

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

	s.logger.Infof("DeleteUser success")

	return &pb.DeleteUserResponse{
		Id: int64(id),
	}, nil
}

func (s *Service) GetUsers(empty *emptypb.Empty, stream pb.UserService_GetUsersServer) error {
	users, err := s.repo.GetUsers(context.Background())
	if err != nil {
		s.logger.Errorf("failed to GetUsers err: %v", err)
		return fmt.Errorf("GetUsers err: %v", err)
	}

	for _, u := range *users {
		user := pb.User{
			Id:          int64(u.ID),
			RoleId:      int64(u.RoleID),
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			Email:       u.Email,
			Phone:       u.Phone,
			Login:       u.Login,
			Password:    u.Password,
			IsConfirmed: u.IsConfirmed,
		}
		if err = stream.Send(&pb.GetUsersResponse{
			Result: &user,
		}); err != nil {
			return status.Errorf(codes.Internal, "fetch: unexpected stream: %v", err)
		}
	}

	s.logger.Infof("GetUsers success")

	return nil
}
