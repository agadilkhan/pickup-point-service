package transport

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/admin/config"
	pb "github.com/agadilkhan/pickup-point-service/pkg/protobuf/userservice/gw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type UserGrpcTransport struct {
	config config.UserGrpcTransport
	client pb.UserServiceClient
}

func NewUserGrpcTransport(config config.UserGrpcTransport) *UserGrpcTransport {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, _ := grpc.Dial(config.Host, opts...)

	client := pb.NewUserServiceClient(conn)

	return &UserGrpcTransport{
		config: config,
		client: client,
	}
}

type UpdateUserRequest struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Login     string
	Password  string
}

func (t *UserGrpcTransport) GetUserByID(ctx context.Context, id int) (*pb.User, error) {
	resp, err := t.client.GetUserByID(ctx, &pb.GetUserByIDRequest{
		Id: int64(id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to GetUserByID err: %v", err)
	}

	return resp.Result, nil
}

func (t *UserGrpcTransport) UpdateUser(ctx context.Context, request UpdateUserRequest) (*pb.User, error) {
	resp, err := t.client.UpdateUser(ctx, &pb.UpdateUserRequest{
		Request: &pb.User{
			Id:        int64(request.ID),
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Phone:     request.Phone,
			Login:     request.Login,
			Password:  request.Password,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to UpdateUser err: %v", err)
	}

	return resp.Result, nil
}

func (t *UserGrpcTransport) DeleteUser(ctx context.Context, id int) (*pb.DeleteUserResponse, error) {
	resp, err := t.client.DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: int64(id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to DeleteUser err: %v", err)
	}

	return &pb.DeleteUserResponse{
		Id: resp.Id,
	}, nil
}

func (t *UserGrpcTransport) GetUsers(ctx context.Context) (*[]pb.GetUsersResponse, error) {
	resp, err := t.client.GetUsers(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to GetUsers err: %v", err)
	}

	var users []pb.GetUsersResponse
	for {
		res, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				return &users, nil
			}
		}

		user := pb.GetUsersResponse{
			Result: &pb.User{
				Id:          res.Result.Id,
				RoleId:      res.Result.RoleId,
				FirstName:   res.Result.FirstName,
				LastName:    res.Result.LastName,
				Email:       res.Result.Email,
				Phone:       res.Result.Phone,
				Login:       res.Result.Login,
				Password:    res.Result.Password,
				IsConfirmed: res.Result.IsConfirmed,
			},
		}

		users = append(users, user)
	}
}
