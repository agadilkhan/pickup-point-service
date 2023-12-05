package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/pkg/pagination"
)

func (s *Service) GetOrders(ctx context.Context, sortOptions pagination.SortOptions, filterOptions pagination.FilterOptions) (*[]entity.Order, error) {
	if sortOptions.SortBy == "" {
		sortOptions.SortBy = "created_at"
	}
	if sortOptions.SortOrder == "" {
		sortOptions.SortOrder = "desc"
	}

	orders, err := s.repo.GetOrders(ctx, sortOptions, filterOptions)
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
