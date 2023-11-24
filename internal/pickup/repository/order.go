package repository

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
)

func (r *Repo) GetOrderByCode(ctx context.Context, code string) (*entity.Order, error) {
	var order entity.Order

	res := r.replica.DB.WithContext(ctx).Where("code = ?", code).First(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	var items []entity.OrderItem

	res = r.replica.DB.WithContext(ctx).Where("order_id = ?", order.ID).Find(&items)
	if res.Error != nil {
		return nil, res.Error
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

	order.OrderItems = items
	order.Customer = customer
	order.Company = company
	order.Point = point

	return &order, nil
}

func (r *Repo) GetOrders(ctx context.Context, sort, direction string) (*[]entity.Order, error) {
	var result []entity.Order
	var orders []entity.Order

	res := r.replica.DB.WithContext(ctx).Where("status != ?", entity.OrderStatusGiven).Order(fmt.Sprintf("%s%s", sort, direction)).Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, order := range orders {
		var items []entity.OrderItem
		res = r.replica.DB.WithContext(ctx).Where("order_id = ?", order.ID).Find(&items)
		if res.Error != nil {
			return nil, res.Error
		}

		var resItems []entity.OrderItem

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

		result = append(result, order)
	}

	return &result, nil
}

func (r *Repo) UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	return nil, nil
}

func (r *Repo) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	res := r.main.DB.WithContext(ctx).Create(order)
	if res.Error != nil {
		return 0, res.Error
	}

	return order.ID, nil
}
