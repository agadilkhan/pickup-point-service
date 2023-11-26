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

func (s *Service) GetWarehouseOrders(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error) {
	orderWarehouses, err := s.repo.GetWarehouseOrders(ctx, warehouseID)
	if err != nil {
		return nil, fmt.Errorf("failed to GetWarehouseOrders err: %v", err)
	}

	return orderWarehouses, nil
}

func (s *Service) GetAllWarehouses(ctx context.Context) (*[]entity.Warehouse, error) {
	warehouses, err := s.repo.GetAllWarehouses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetWarehouses err: %v", err)
	}

	return warehouses, nil
}
