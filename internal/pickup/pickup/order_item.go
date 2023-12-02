package pickup

import (
	"context"
	"fmt"
)

func (s *Service) ReceiveItem(ctx context.Context, orderCode string, productID int) error {
	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	for _, item := range order.OrderItems {
		if productID == item.ProductID && !item.IsAccept {
			item.IsAccept = true
			_, err = s.repo.UpdateItem(ctx, &item)
			if err != nil {
				return fmt.Errorf("failed to UpdateItem err: %v", err)
			}

			return nil
		}
	}

	return fmt.Errorf("failed to receive item")
}

func (s *Service) RefundItem(ctx context.Context, orderCode string, request RefundItemRequest) error {
	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	for _, item := range order.OrderItems {
		if item.ProductID == request.ProductID && item.NumOfRefund == 0 && item.Quantity >= request.Quantity {
			item.NumOfRefund = request.Quantity
			_, err = s.repo.UpdateItem(ctx, &item)
			if err != nil {
				return fmt.Errorf("failed to UpdateItem err: %v", err)
			}

			return nil
		}
	}

	return fmt.Errorf("failed to refund item")
}
