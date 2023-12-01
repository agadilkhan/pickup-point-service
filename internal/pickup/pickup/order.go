package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"math/rand"
	"reflect"
	"sort"
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

	if order.PaymentStatus == entity.PaymentStatusNotPaid {
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

	_, err = s.repo.CreateTransaction(ctx, &transaction)

	if err != nil {
		return fmt.Errorf("failed to CreateTransaction err: %v", err)
	}

	err = s.repo.DeleteWarehouseOrderByOrderID(ctx, order.ID)
	if err != nil {
		return fmt.Errorf("failed to DeleteWarehouseOrderByOrderID err: %v", err)
	}

	return nil

}

func (s *Service) ReceiveOrder(ctx context.Context, request ReceiveOrderRequest) (int, error) {
	ctxUser, ok := ctx.Value("user_id").(float64)
	if !ok {
		return -1, fmt.Errorf("failed to parse user_id from context")
	}

	userID := int(ctxUser)

	order, err := s.GetOrderByCode(ctx, request.OrderCode)
	if err != nil {
		return -1, fmt.Errorf("failed to GetOrderByCode")
	}

	if order.Status != entity.OrderStatusProcessing {
		return -1, fmt.Errorf("order not ready to receive")
	}

	// comparing products by id
	var arr1, arr2 []int

	for _, item := range request.Items {
		arr1 = append(arr1, item.ProductID)
	}
	for _, item := range order.OrderItems {
		arr2 = append(arr2, item.ProductID)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	if len(arr1) != len(arr2) {
		return -1, fmt.Errorf("wrong quantity of products")
	}
	if !reflect.DeepEqual(arr1, arr2) {
		return -1, fmt.Errorf("wrong products")
	}

	// checking free place in warehouse
	warehouse, err := s.GetWarehouseByID(ctx, request.WarehouseID)
	if err != nil {
		return -1, fmt.Errorf("failed to GetWarehouseByID err: %v", err)
	}

	warehouseOrders, err := s.repo.GetWarehouseOrdersByWarehouseID(ctx, warehouse.ID)
	if err != nil {
		return -1, fmt.Errorf("failed to GetWarehouseOrders err: %v", err)
	}
	if len(*warehouseOrders) == warehouse.NumOfPlaces {
		return -1, fmt.Errorf("not free place")
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

	order.Status = entity.OrderStatusDelivered

	warehouseOrder := entity.WarehouseOrder{
		WarehouseID: warehouse.ID,
		OrderID:     order.ID,
		PlaceNum:    placeNum,
		Order:       *order,
	}

	_, err = s.repo.CreateWarehouseOrder(ctx, &warehouseOrder)
	if err != nil {
		return -1, fmt.Errorf("failed to CreateWarehouseOrder err: %v", err)
	}

	transaction := entity.Transaction{
		UserID:          userID,
		OrderID:         order.ID,
		TransactionType: entity.TransactionTypeReceive,
		Order:           *order,
	}

	_, err = s.repo.CreateTransaction(ctx, &transaction)
	if err != nil {
		return -1, fmt.Errorf("failed to CreateTransaction err: %v", err)
	}

	return placeNum, nil
}
