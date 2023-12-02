package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"math/rand"
	"sync"
)

func (s *Service) GetOrders(ctx context.Context, query GetOrdersQuery) (*[]entity.Order, error) {
	if query.Sort == "" {
		query.Sort = "created_at"
	}
	if query.Direction == "" {
		query.Direction = "desc"
	}

	orders, err := s.repo.GetOrders(ctx, query.Sort, query.Direction)
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

func (s *Service) CreateOrder(ctx context.Context, request CreateOrderRequest) (int, error) {
	var orderItems []entity.OrderItem

	// converting request struct for order_item
	for _, item := range request.Items {
		product, err := s.GetProduct(ctx, item.ProductID)
		if err != nil {
			return 0, fmt.Errorf("failed to GetProductByID err: %v", err)
		}

		subTotal := product.Price * float64(item.Quantity)

		orderItem := entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			SubTotal:  subTotal,
			Product:   *product,
		}

		orderItems = append(orderItems, orderItem)
	}

	// calculating total amount of order
	var totalAmount float64
	for _, orderItem := range orderItems {
		totalAmount += orderItem.SubTotal
	}

	code := s.GenerateOrderCode(request.CustomerID)

	order := entity.Order{
		CustomerID:  request.CustomerID,
		CompanyID:   request.CompanyID,
		PointID:     request.PointID,
		Code:        code,
		Status:      entity.OrderStatusProcessing,
		IsPaid:      request.IsPaid,
		TotalAmount: totalAmount,
		OrderItems:  orderItems,
	}

	id, err := s.repo.CreateOrder(ctx, &order)
	if err != nil {
		return 0, fmt.Errorf("failed to CreateOrder err: %v", err)
	}

	return id, nil
}

func (s *Service) DeleteOrder(ctx context.Context, orderCode string) (string, error) {
	_, err := s.repo.DeleteOrder(ctx, orderCode)
	if err != nil {
		return "", fmt.Errorf("failed to DeleteOrder err: %v", err)
	}

	return orderCode, nil
}

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

	if order.Status != entity.OrderStatusProcessing {
		return fmt.Errorf("order is not ready to receive")
	}

	for _, item := range order.OrderItems {
		if !item.IsAccept {
			return fmt.Errorf("item has not been accept")
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
	if order.Status == entity.OrderStatusCancelled {
		return fmt.Errorf("order is already cancelled")
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
	if order.Status == entity.OrderStatusRefund {
		return fmt.Errorf("order is already refund")
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
