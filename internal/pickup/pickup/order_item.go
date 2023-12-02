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
	errCh := make(chan error, len(order.OrderItems))

	found := false

	for _, item := range order.OrderItems {
		wg.Add(1)
		go func(orderItem entity.OrderItem) {
			defer wg.Done()

			if productID == orderItem.ProductID && !orderItem.IsAccept {
				orderItem.IsAccept = true
				found = true
				_, err = s.repo.UpdateItem(ctx, &orderItem)
				if err != nil {
					errCh <- err
				}
			}
		}(item)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err = range errCh {
		if err != nil {
			return err
		}
	}

	if !found {
		return fmt.Errorf("the item has not been found")
	}

	return nil
}

func (s *Service) RefundItem(ctx context.Context, orderCode string, request RefundItemRequest) error {
	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	found := false

	var wg sync.WaitGroup
	errCh := make(chan error, len(order.OrderItems))

	for _, item := range order.OrderItems {
		wg.Add(1)
		go func(orderItem entity.OrderItem) {
			defer wg.Done()
			if orderItem.ProductID == request.ProductID && orderItem.Quantity >= request.Quantity {
				orderItem.NumOfRefund = request.Quantity
				found = true
				_, err = s.repo.UpdateItem(ctx, &orderItem)
				if err != nil {
					errCh <- err
				}
			}
		}(item)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err = range errCh {
		if err != nil {
			return err
		}
	}

	if !found {
		return fmt.Errorf("the item has not been found")
	}

	return nil
}
