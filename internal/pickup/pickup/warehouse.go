package pickup

//func (s *Service) GetWarehouseOrders(ctx context.Context, warehouseID int) (*[]entity.WarehouseOrder, error) {
//	warehouseOrders, err := s.repo.GetWarehouseOrders(ctx, warehouseID)
//	if err != nil {
//		return nil, err
//	}
//
//	return warehouseOrders, nil
//}
//
//func (s *Service) GetWarehouse(ctx context.Context, pointID int) (*entity.Warehouse, error) {
//	warehouse, err := s.repo.GetWarehouse(ctx, pointID)
//	if err != nil {
//		return nil, err
//	}
//
//	return warehouse, nil
//}
//
//func (s *Service) DeleteWarehouseOrder(ctx context.Context, orderID int) error {
//	err := s.repo.DeleteWarehouseOrder(ctx, orderID)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *Service) CreateWarehouseOrder(ctx context.Context, order *entity.WarehouseOrder) (int, error) {
//	id, err := s.repo.CreateWarehouseOrder(ctx, order)
//	if err != nil {
//		return 0, err
//	}
//
//	return id, nil
//}
