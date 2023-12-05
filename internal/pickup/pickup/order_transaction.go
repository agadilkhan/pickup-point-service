package pickup

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) PickupOrder(ctx context.Context, code string) error {
	ctxUser, ok := ctx.Value("user_id").(float64)
	if !ok {
		return fmt.Errorf("cannot parse user_id from context")
	}

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

	if !order.IsPaid {
		return fmt.Errorf("order not paid")
	}

	userID := int(ctxUser)

	transaction := entity.Transaction{
		UserID:          userID,
		OrderID:         order.ID,
		TransactionType: entity.TransactionTypePickup,
		Order:           *order,
	}

	transaction.Order.Status = entity.OrderStatusGiven

	_, err = s.CreateTransaction(ctx, &transaction)
	if err != nil {
		return fmt.Errorf("failed to CreateTransaction err: %v", err)
	}

	return nil

}

func (s *Service) ReceiveOrder(ctx context.Context, orderCode string) error {
	ctxUser, ok := ctx.Value("user_id").(float64)
	if !ok {
		return fmt.Errorf("failed to parse user_id from context")
	}

	userID := int(ctxUser)

	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(order.OrderItems))

	for _, item := range order.OrderItems {
		wg.Add(1)
		go func(orderItem entity.OrderItem) {
			defer wg.Done()
			if !orderItem.IsAccept {
				errCh <- fmt.Errorf("item has not been accept")
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

	order.Status = entity.OrderStatusDelivered

	transaction := entity.Transaction{
		UserID:          userID,
		OrderID:         order.ID,
		TransactionType: entity.TransactionTypeReceive,
		Order:           *order,
	}

	_, err = s.repo.CreateTransaction(ctx, &transaction)
	if err != nil {
		return fmt.Errorf("failed to CreateTransaction err: %v", err)
	}

	return nil
}

func (s *Service) CancelOrder(ctx context.Context, orderCode string) error {
	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	order.Status = entity.OrderStatusCancelled

	_, err = s.repo.UpdateOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to UpdateOrder err: %v", err)
	}

	return nil
}

func (s *Service) RefundOrder(ctx context.Context, orderCode string) error {
	ctxUser, ok := ctx.Value("user_id").(float64)
	if !ok {
		return fmt.Errorf("failed to parse user_id from context")
	}

	userID := int(ctxUser)

	order, err := s.GetOrderByCode(ctx, orderCode)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(order.OrderItems))

	for _, item := range order.OrderItems {
		wg.Add(1)
		go func(orderItem entity.OrderItem) {
			defer wg.Done()

			orderItem.NumOfRefund = orderItem.Quantity
			_, err = s.repo.UpdateItem(ctx, &orderItem)
			if err != nil {
				errCh <- err
			}
		}(item)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err = range errCh {
		if err != nil {
			return fmt.Errorf("failed to UpdateItem err: %v", err)
		}
	}

	order.Status = entity.OrderStatusRefund

	transaction := entity.Transaction{
		UserID:          userID,
		OrderID:         order.ID,
		TransactionType: entity.TransactionTypeRefund,
		Order:           *order,
	}

	_, err = s.CreateTransaction(ctx, &transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GenerateOrderCode(customerID int) string {
	randNum1 := rand.Intn(10)
	randNum2 := rand.Intn(10)
	randNum3 := rand.Intn(10)
	randNum4 := rand.Intn(10)

	// generating unique code for order
	code := fmt.Sprintf("%d * %d%d%d%d", customerID, randNum1, randNum2, randNum3, randNum4)

	return code
}
