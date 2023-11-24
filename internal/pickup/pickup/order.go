package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetOrders(ctx context.Context, request GetAllOrdersRequest) (*[]entity.Order, error) {
	if request.Sort == "" {
		request.Sort = "id"
	}
	if request.Direction == "" {
		request.Direction = "asc"
	}

	orders, err := s.repo.GetOrders(ctx, request.Sort, request.Direction)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllOrders err: %v", err)
	}

	return orders, nil
}

func (s *Service) GetOrderByCode(ctx context.Context, code string) (*entity.Order, error) {
	order, err := s.repo.GetOrderByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to GetOrderByCode err: %v", err)
	}

	return order, nil
}

func (s *Service) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	return 0, nil
}
