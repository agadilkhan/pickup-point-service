package pickup

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/transport"
)

type Service struct {
	repo     repository.Repository
	userGrpc *transport.UserGrpcTransport
}

func NewPickupService(repo repository.Repository, userGrpc *transport.UserGrpcTransport) *Service {
	return &Service{
		repo:     repo,
		userGrpc: userGrpc,
	}
}
