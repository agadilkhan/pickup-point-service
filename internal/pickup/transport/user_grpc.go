package transport

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
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
