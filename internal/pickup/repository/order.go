package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/pkg/pagination"
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

	var resItems = make([]entity.OrderItem, 0)
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

	order.OrderItems = items
	order.Customer = customer
	order.Company = company
	order.Point = point
	order.OrderItems = resItems

	return &order, nil
}

func (r *Repo) GetOrders(ctx context.Context, sortOptions pagination.SortOptions, filterOptions pagination.FilterOptions) (*[]entity.Order, error) {
	var result = make([]entity.Order, 0)
	var orders []entity.Order

	for _, field := range filterOptions.Fields {
		res := r.replica.DB.WithContext(ctx).Where(field.GetQuery())
		if res.Error != nil {
			return nil, res.Error
		}
	}

	res := r.replica.DB.WithContext(ctx).Order(sortOptions.GetOrderBy()).Find(&orders)
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
	res := r.main.DB.WithContext(ctx).Where("id = ?", order.ID).Updates(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	return order, nil
}

func (r *Repo) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	res := r.main.DB.WithContext(ctx).Create(&order)
	if res.Error != nil {
		return 0, res.Error
	}

	return order.ID, nil
}

func (r *Repo) DeleteOrder(ctx context.Context, orderCode string) (string, error) {
	res := r.main.DB.WithContext(ctx).Where("code = ?", orderCode).Delete(&entity.Order{})
	if res.Error != nil {
		return "", res.Error
	}

	return orderCode, nil
}
