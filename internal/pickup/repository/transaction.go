package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"gorm.io/gorm"
)

func (r *Repo) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (int, error) {
	err := r.main.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&transaction)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Where("id = ?", transaction.Order.ID).Updates(&transaction.Order)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return transaction.ID, nil
}

func (r *Repo) GetTransactions(ctx context.Context, userID int, transactionType string) (*[]entity.Transaction, error) {
	var resTransactions = make([]entity.Transaction, 0)
	var transactions []entity.Transaction

	var res *gorm.DB

	if transactionType == "" {
		res = r.replica.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&transactions)
	} else {
		res = r.replica.DB.WithContext(ctx).Where("user_id = ? AND transaction_type = ?", userID, transactionType).Find(&transactions)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	for _, transaction := range transactions {
		var order entity.Order
		res = r.replica.DB.WithContext(ctx).Where("id = ?", transaction.OrderID).Find(&order)
		if res.Error != nil {
			return nil, res.Error
		}

		var resItems []entity.OrderItem
		var items []entity.OrderItem

		res = r.replica.DB.WithContext(ctx).Where("order_id = ?", order.ID).Find(&items)
		if res.Error != nil {
			return nil, res.Error
		}

		for _, item := range items {
			var product entity.Product
			res = r.replica.DB.WithContext(ctx).Where("id = ?", item.ProductID).First(&product)
			if res.Error != nil {
				return nil, res.Error
			}

			item.Product = product
			resItems = append(resItems, item)
		}

		var customer entity.Customer
		res = r.replica.DB.WithContext(ctx).Where("id = ?", order.CustomerID).First(&customer)
		if res.Error != nil {
			return nil, res.Error
		}

		var company entity.Company
		res = r.replica.DB.WithContext(ctx).Where("id = ?", order.CompanyID).First(&company)
		if res.Error != nil {
			return nil, res.Error
		}

		var point entity.PickupPoint
		res = r.replica.DB.WithContext(ctx).Where("id = ?", order.PointID).First(&point)
		if res.Error != nil {
			return nil, res.Error
		}

		order.OrderItems = resItems
		order.Customer = customer
		order.Company = company
		order.Point = point

		transaction.Order = order

		resTransactions = append(resTransactions, transaction)
	}

	return &resTransactions, nil
}
