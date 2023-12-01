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

func (r *Repo) GetWarehouseOrdersByWarehouseID(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error) {
	var warehouseOrders []entity.WarehouseOrder

	res := r.replica.DB.WithContext(ctx).Where("warehouse_id = ?", warehouseID).Find(&warehouseOrders)
	if res.Error != nil {
		return nil, res.Error
	}

	return &warehouseOrders, nil
}

func (r *Repo) CreateWarehouseOrder(ctx context.Context, warehouseOrder *entity.WarehouseOrder) (int, error) {
	err := r.main.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.WithContext(ctx).Create(&warehouseOrder)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return warehouseOrder.ID, nil
}

func (r *Repo) GetWarehouses(ctx context.Context) (*[]entity.Warehouse, error) {
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

func (r *Repo) DeleteWarehouseOrderByOrderID(ctx context.Context, orderID int) error {
	res := r.main.DB.WithContext(ctx).Where("order_id = ?", orderID).Delete(entity.WarehouseOrder{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
