package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"math/rand"
	"sort"
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

func (s *Service) CreateOrder(ctx context.Context, request CreateOrderRequest) (int, error) {
	var orderItems []entity.OrderItem

	// converting request struct for order_item
	for _, item := range request.Items {
		product, err := s.GetProductByID(ctx, item.ProductID)
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

	randNum1 := rand.Intn(10)
	randNum2 := rand.Intn(10)
	randNum3 := rand.Intn(10)
	randNum4 := rand.Intn(10)

	// generating unique code for order
	code := fmt.Sprintf("%d * %d%d%d%d", request.CustomerID, randNum1, randNum2, randNum3, randNum4)

	order := entity.Order{
		CustomerID:    request.CustomerID,
		CompanyID:     request.CompanyID,
		PointID:       request.PointID,
		Code:          code,
		Status:        entity.OrderStatusProcessing,
		PaymentStatus: request.PaymentStatus,
		TotalAmount:   totalAmount,
		OrderItems:    orderItems,
	}

	id, err := s.repo.CreateOrder(ctx, &order)
	if err != nil {
		return 0, fmt.Errorf("failed to CreateOrder err: %v", err)
	}

	return id, nil
}

func (s *Service) ReceiveOrder(ctx context.Context, request ReceiveOrderRequest) (int, error) {
	order, err := s.GetOrderByCode(ctx, request.OrderCode)
	if err != nil {
		return 0, fmt.Errorf("failed to GetOrderByCode")
	}

	if order.Status != entity.OrderStatusProcessing {
		return 0, fmt.Errorf("order not ready to receive")
	}

	// comparing products by id
	var arr1, arr2 []int

	for _, item := range request.Items {
		arr1 = append(arr1, item.ProductID)
	}

	for _, item := range order.OrderItems {
		arr2 = append(arr2, item.ProductID)
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	if len(arr1) != len(arr2) {
		return 0, fmt.Errorf("wrong quantity of products")
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return 0, fmt.Errorf("wrong product")
		}
	}

	// checking free place in warehouse
	warehouse, err := s.repo.GetWarehouseByID(ctx, request.WarehouseID)
	if err != nil {
		return 0, fmt.Errorf("failed to GetWarehouseByID err: %v", err)
	}

	if warehouse.NumOfFreePlaces == 0 {
		return 0, fmt.Errorf("not free places")
	}

	warehouseOrders, err := s.repo.GetWarehouseOrders(ctx, request.WarehouseID)
	if err != nil {
		return 0, fmt.Errorf("failed to GetWarehouseOrders err: %v", err)
	}

	// selecting place from warehouse for order
	var existingPlaces []int
	for _, o := range *warehouseOrders {
		existingPlaces = append(existingPlaces, o.PlaceNum)
	}

	var placeNum int
	for {
		placeNum = rand.Intn(warehouse.NumOfPlaces)

		found := false
		for _, num := range existingPlaces {
			if placeNum == num {
				found = true
				break
			}
		}

		if !found {
			break
		}
	}
	warehouse.NumOfFreePlaces -= 1

	warehouseOrder := entity.WarehouseOrder{
		WarehouseID: warehouse.ID,
		OrderID:     order.ID,
		PlaceNum:    placeNum,
		Order:       *order,
		Warehouse:   *warehouse,
	}

	err = s.repo.CreateWarehouseOrder(ctx, &warehouseOrder)
	if err != nil {
		return 0, fmt.Errorf("failed to ReceiveOrder err: %v", err)
	}

	return placeNum, nil
}
