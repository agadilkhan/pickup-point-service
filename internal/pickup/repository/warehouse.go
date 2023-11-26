package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"gorm.io/gorm"
)

func (r *Repo) GetWarehouseByID(ctx context.Context, id int) (*entity.Warehouse, error) {
	var warehouse entity.Warehouse

	res := r.replica.DB.WithContext(ctx).Where("id = ?", id).First(&warehouse)
	if res.Error != nil {
		return nil, res.Error
	}

	var point entity.PickupPoint

	res = r.replica.DB.WithContext(ctx).Where("id = ?", warehouse.PointID).First(&point)
	if res.Error != nil {
		return nil, res.Error
	}

	warehouse.Point = point

	return &warehouse, nil
}

func (r *Repo) GetWarehouseOrders(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error) {
	var warehouseOrders []entity.WarehouseOrder

	res := r.replica.DB.WithContext(ctx).Where("warehouse_id = ?", warehouseID).Find(&warehouseOrders)
	if res.Error != nil {
		return nil, res.Error
	}

	return &warehouseOrders, nil
}

func (r *Repo) CreateWarehouseOrder(ctx context.Context, warehouseOrder *entity.WarehouseOrder) error {
	err := r.main.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.WithContext(ctx).Create(warehouseOrder)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Model(&warehouseOrder.Order).WithContext(ctx).Updates(entity.Order{
			Status: entity.OrderStatusDelivered,
		})
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Model(&warehouseOrder.Warehouse).Updates(entity.Warehouse{
			NumOfFreePlaces: warehouseOrder.Warehouse.NumOfFreePlaces,
		})
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAllWarehouses(ctx context.Context) (*[]entity.Warehouse, error) {
	var result = make([]entity.Warehouse, 0)
	var warehouses []entity.Warehouse

	res := r.replica.DB.WithContext(ctx).Find(&warehouses)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, warehouse := range warehouses {
		var point entity.PickupPoint

		res = r.replica.DB.WithContext(ctx).Where("id = ?", warehouse.PointID).First(&point)
		if res.Error != nil {
			return nil, res.Error
		}

		warehouse.Point = point

		result = append(result, warehouse)
	}

	return &result, nil
}
