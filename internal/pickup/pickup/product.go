package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetProductByID(ctx context.Context, id int) (*entity.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetProductByID err: %v", err)
	}

	return product, nil
}

func (s *Service) GetProducts(ctx context.Context) (*[]entity.Product, error) {
	products, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetProducts err: %v", err)
	}

	return products, nil
}
