package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"gorm.io/gorm"
)

func (r *Repo) GetAllPickupPoints(ctx context.Context) (*[]entity.PickupPoint, error) {
	var points []entity.PickupPoint

	res := r.replica.DB.WithContext(ctx).Find(&points)
	if res.Error != nil {
		return nil, res.Error
	}

	return &points, nil
}

func (r *Repo) GetPickupPointByID(ctx context.Context, id int) (*entity.PickupPoint, error) {
	var point entity.PickupPoint

	res := r.replica.DB.WithContext(ctx).Where("id = ?", id).First(&point)
	if res.Error != nil {
		return nil, res.Error
	}

	return &point, nil
}

func (r *Repo) CreatePickupOrder(ctx context.Context, pickup *entity.PickupOrder) (int, error) {
	err := r.main.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.WithContext(ctx).Create(&pickup)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Model(&pickup.Order).WithContext(ctx).Updates(entity.Order{
			Status: entity.OrderStatusGiven,
		})
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		var warehouseOrder entity.OrderWarehouse

		res = tx.WithContext(ctx).Where("order_id = ?", pickup.OrderID).First(&warehouseOrder)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		var warehouse entity.Warehouse

		res = tx.WithContext(ctx).Where("id = ?", warehouseOrder.WarehouseID).First(&warehouse)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Model(&warehouse).WithContext(ctx).Updates(entity.Warehouse{
			NumOfFreePlaces: warehouseOrder.Warehouse.NumOfFreePlaces + 1,
		})
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.WithContext(ctx).Delete(&warehouseOrder)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return pickup.ID, nil
}

func (r *Repo) GetPickupOrders(ctx context.Context, userID int) (*[]entity.PickupOrder, error) {
	var result []entity.PickupOrder
	var pickupOrders []entity.PickupOrder

	res := r.replica.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&pickupOrders)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, pickupOrder := range pickupOrders {
		var order entity.Order
		res = r.replica.DB.WithContext(ctx).Where("id = ?", pickupOrder.OrderID).First(&order)
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

		pickupOrder.Order = order

		result = append(result, pickupOrder)
	}

	return &result, nil
}

func (r *Repo) GetPickupOrderByID(ctx context.Context, userID, pickupOrderID int) (*entity.PickupOrder, error) {
	var pickupOrder entity.PickupOrder
	res := r.replica.DB.WithContext(ctx).Where("id = ? AND user_id = ?", pickupOrderID, userID).First(&pickupOrder)
	if res.Error != nil {
		return nil, res.Error
	}

	var order entity.Order
	res = r.replica.DB.WithContext(ctx).Where("id = ?", pickupOrder.OrderID).First(&order)
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

	pickupOrder.Order = order

	return &pickupOrder, nil
}
