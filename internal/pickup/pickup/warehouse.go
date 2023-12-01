package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetWarehouseByID(ctx context.Context, id int) (*entity.Warehouse, error) {
	warehouse, err := s.repo.GetWarehouseByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetWarehouseByID err: %v", err)
	}

	return warehouse, nil
}

func (s *Service) GetWarehouses(ctx context.Context) (*[]entity.Warehouse, error) {
	warehouses, err := s.repo.GetWarehouses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetWarehouses err: %v", err)
	}

	return warehouses, nil
}
