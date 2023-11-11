package order

import (
	"github.com/agadilkhan/pickup-point-service/internal/order/config"
	"github.com/agadilkhan/pickup-point-service/internal/order/repository"
)

type Deps struct {
	repo *repository.Repository
	cfg  *config.Config
}

func NewDeps(repo *repository.Repository, cfg *config.Config) Deps {
	return Deps{
		repo: repo,
		cfg:  cfg,
	}
}

type Service struct {
	Order
}

func NewService(deps Deps) *Service {
	return &Service{
		Order: NewOrderService(deps.repo.Order),
	}
}
