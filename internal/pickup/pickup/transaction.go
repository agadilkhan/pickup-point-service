package pickup

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (s *Service) GetTransactions(ctx context.Context, userID int, query GetTransactionsQuery) (*[]entity.Transaction, error) {
	transactions, err := s.repo.GetTransactions(ctx, userID, query.TransactionType)
	if err != nil {
		return nil, fmt.Errorf("failed to GetTransactions err: %v", err)
	}

	return transactions, nil
}
