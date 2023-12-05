package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/pkg/pagination"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	product, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) GetProducts(ctx context.Context, searchOptions pagination.SearchOptions) (*[]entity.Product, error) {
	products, err := s.repo.GetProducts(ctx, searchOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to GetProducts err: %v", err)
	}

	return products, nil
}
