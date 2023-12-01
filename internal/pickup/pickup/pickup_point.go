package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error) {
	points, err := s.repo.GetPickupPoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllPickupPoints err: %v", err)
	}

	return points, nil
}

func (s *Service) GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error) {
	point, err := s.repo.GetPickupPointByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetPickupOrderByID err: %v", err)
	}

	return point, nil
}
