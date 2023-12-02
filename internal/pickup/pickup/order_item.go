package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"sync"
)

func (s *Service) ReceiveItem(ctx context.Context, orderCode string, productID int) error {
	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 0)

	for _, item := range order.OrderItems {
		wg.Add(1)
		go func(orderItem entity.OrderItem) {
			if productID == orderItem.ProductID && !orderItem.IsAccept {
				orderItem.IsAccept = true
				_, err = s.repo.UpdateItem(ctx, &orderItem)
				if err != nil {
					errCh <- err
				}
			}
		}(item)
	}

	go func() {
		wg.Done()
		close(errCh)
	}()

	for err = range errCh {
		if err != nil {
			return err
		}
	}

	return fmt.Errorf("failed to receive item")
}

func (s *Service) RefundItem(ctx context.Context, orderCode string, request RefundItemRequest) error {
	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 0)

	for _, item := range order.OrderItems {
		wg.Add(1)
		go func(orderItem entity.OrderItem) {
			if orderItem.ProductID == request.ProductID && orderItem.Quantity >= request.Quantity {
				orderItem.NumOfRefund = request.Quantity
				_, err = s.repo.UpdateItem(ctx, &orderItem)
				if err != nil {
					errCh <- err
				}
			}
		}(item)
	}

	go func() {
		wg.Done()
		close(errCh)
	}()

	for err = range errCh {
		if err != nil {
			return err
		}
	}

	return fmt.Errorf("failed to refund item")
}
