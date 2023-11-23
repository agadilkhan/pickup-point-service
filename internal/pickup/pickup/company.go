package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetCompanyByID(ctx context.Context, id int) (*entity.Company, error) {
	company, err := s.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetCompanyByID err: %v", err)
	}

	return company, nil
}
