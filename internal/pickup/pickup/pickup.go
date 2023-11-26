package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) Pickup(ctx context.Context, code string) error {
	order, err := s.GetOrderByCode(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to GetOrderByCode err: %v", err)
	}

	if order.Status != entity.OrderStatusReady {
		if order.Status == entity.OrderStatusGiven {
			return fmt.Errorf("order is already given")
		}
		return fmt.Errorf("order not ready to pickup")
	}

	if order.PaymentStatus == entity.PaymentStatusNotPaid {
		return fmt.Errorf("order not paid")
	}

	val, ok := ctx.Value("user_id").(float64)
	if !ok {
		return fmt.Errorf("cannot convert to float64")
	}

	userID := int(val)

	orderPickup := entity.PickupOrder{
		UserID:  userID,
		OrderID: order.ID,
		Order:   *order,
	}

	_, err = s.repo.CreatePickupOrder(ctx, &orderPickup)
	if err != nil {
		return fmt.Errorf("failed to CreateOrderPickup err: %v", err)
	}

	return nil
}

func (s *Service) GetPickupOrders(ctx context.Context, userID int) (*[]entity.PickupOrder, error) {
	pickupOrders, err := s.repo.GetPickupOrders(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to GetOrderPickups err: %v", err)
	}

	return pickupOrders, nil
}

func (s *Service) GetPickupOrderByID(ctx context.Context, request GetPickupOrderByIDRequest) (*entity.PickupOrder, error) {
	pickupOrder, err := s.repo.GetPickupOrderByID(ctx, request.UserID, request.PickupOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to GetPickupOrderByID err: %v", err)
	}

	return pickupOrder, nil
}

func (s *Service) GetAllPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error) {
	points, err := s.repo.GetAllPickupPoints(ctx)
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
