package admin

import "github.com/agadilkhan/pickup-point-service/internal/admin/transport"

type Service struct {
	userGrpcTransport *transport.UserGrpcTransport
}

func NewService(userGrpcTransport *transport.UserGrpcTransport) *Service {
	return &Service{
		userGrpcTransport: userGrpcTransport,
	}
}
