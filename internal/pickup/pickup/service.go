package pickup

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository"
)

type Service struct {
	repo repository.Repository
}

func NewPickupService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
