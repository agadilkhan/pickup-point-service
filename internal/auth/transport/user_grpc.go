package transport

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	pb "github.com/agadilkhan/pickup-point-service/pkg/protobuf/userservice/gw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGrpcTransport struct {
	cfg    config.UserGrpcTransport
	client pb.UserServiceClient
}

func NewUserGrpcTransport(cfg config.UserGrpcTransport) *UserGrpcTransport {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, _ := grpc.Dial(cfg.Host, opts...)

	client := pb.NewUserServiceClient(conn)

	return &UserGrpcTransport{
		cfg:    cfg,
		client: client,
	}
}

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Login     string
	Password  string
}

func (t *UserGrpcTransport) GetUserByLogin(ctx context.Context, login string) (*pb.User, error) {
	resp, err := t.client.GetUserByLogin(ctx, &pb.GetUserByLoginRequest{
		Login: login,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to GetUserByLogin err: %v", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("not found")
	}

	return resp.Result, nil
}

func (t *UserGrpcTransport) CreateUser(ctx context.Context, request CreateUserRequest) (*pb.CreateUserResponse, error) {
	resp, err := t.client.CreateUser(ctx, &pb.CreateUserRequest{
		Request: &pb.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Phone:     request.Phone,
			Login:     request.Login,
			Password:  request.Password,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to CreateUser err: %v", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("not create")
	}

	return &pb.CreateUserResponse{
		Id: resp.Id,
	}, nil
}

func (t *UserGrpcTransport) ConfirmUser(ctx context.Context, email string) error {
	_, err := t.client.ConfirmUser(ctx, &pb.ConfirmUserRequest{
		Email: email,
	})
	if err != nil {
		return fmt.Errorf("failed to ConfirmUser err: %v", err)
	}

	return nil
}
