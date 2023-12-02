package pickup

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	product, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
