package pickup

type GetAllOrdersRequest struct {
	Sort      string
	Direction string
}

type GetPickupOrderByIDRequest struct {
	UserID        int
	PickupOrderID int
}
