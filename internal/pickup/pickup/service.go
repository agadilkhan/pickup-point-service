package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/transport"
)

type Service struct {
	repo     repository.Repository
	userGrpc *transport.UserGrpcTransport
}

func NewPickupService(repo repository.Repository, userGrpc *transport.UserGrpcTransport) *Service {
	return &Service{
		repo:     repo,
		userGrpc: userGrpc,
	}
}

func (s *Service) Pickup(ctx context.Context, code string) error {
	order, err := s.repo.GetOrderByCode(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to GetOrderByCode err: %v", err)
	}

	if order.Status != entity.OrderStatusReady {
		if order.Status == entity.OrderStatusGiven {
			return fmt.Errorf("order is already given")
		}
		return fmt.Errorf("order not ready to pickup")
	}

	val, ok := ctx.Value("user_id").(float64)
	if !ok {
		return fmt.Errorf("cannot convert to float64")
	}

	userID := int(val)

	orderPickup := entity.OrderPickup{
		UserID:  userID,
		OrderID: order.ID,
		Order:   *order,
	}

	_, err = s.repo.CreateOrderPickup(ctx, &orderPickup)
	if err != nil {
		return fmt.Errorf("failed to CreateOrderPickup err: %v", err)
	}

	return nil
}

func (s *Service) GetOrderByCode(ctx context.Context, code string) (*entity.Order, error) {
	order, err := s.repo.GetOrderByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to GetOrderByCode err: %v", err)
	}

	return order, nil
}

func (s *Service) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	return 0, nil
}

func (s *Service) GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error) {
	customer, err := s.repo.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetCustomerByID err: %v", err)
	}

	return customer, nil
}

func (s *Service) GetCompanyByID(ctx context.Context, id int) (*entity.Company, error) {
	company, err := s.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetCompanyByID err: %v", err)
	}

	return company, nil
}

func (s *Service) UserInfo(ctx context.Context, login string) (*entity.User, error) {
	user, err := s.userGrpc.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("failed to GetUserByLogin err: %v", err)
	}

	res := entity.User{
		ID:          int(user.Id),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		Login:       user.Login,
		Password:    user.Password,
		IsConfirmed: user.IsConfirmed,
	}

	return &res, nil
}

func (s *Service) ReceiveOrder(ctx context.Context, code string) error {
	return nil
}
