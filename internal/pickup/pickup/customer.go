package pickup

import (
	"context"
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetCustomerByID(ctx context.Context, id int) (*entity.Customer, error) {
	customer, err := s.repo.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetCustomerByID err: %v", err)
	}

	return customer, nil
}

func (s *Service) GetAllCustomers(ctx context.Context) (*[]entity.Customer, error) {
	customers, err := s.repo.GetAllCustomers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllCustomers err: %v", err)
	}

	return customers, nil
}
