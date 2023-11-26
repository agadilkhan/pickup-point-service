package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (eh *EndpointHandler) initOrderRoutes(api *gin.RouterGroup) {
	orders := api.Group("/orders")
	{
		orders.GET("/", eh.GetOrders)
		orders.POST("/", eh.CreateOrder)
		orders.GET("/:code", eh.GetOrderByCode)
		orders.POST("/:code/pickup", eh.Pickup)
		orders.POST("/:code/receive", eh.ReceiveOrder)
	}
}

func (eh *EndpointHandler) Pickup(ctx *gin.Context) {
	code := ctx.Param("code")

	err := eh.service.Pickup(ctx, code)
	if err != nil {
		eh.logger.Errorf("failed to Pickup err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (eh *EndpointHandler) CreateOrder(ctx *gin.Context) {
	request := struct {
		CustomerID    int                  `json:"customer_id"`
		CompanyID     int                  `json:"company_id"`
		PointID       int                  `json:"point_id"`
		Status        entity.OrderStatus   `json:"status"`
		PaymentStatus entity.PaymentStatus `json:"payment_status"`
		Items         []struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"items"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		eh.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	orderID, err := eh.service.CreateOrder(ctx, request)
	if err != nil {
		eh.logger.Errorf("failed to CreateOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusCreated, orderID)
}

func (eh *EndpointHandler) GetOrders(ctx *gin.Context) {
	sort := ctx.Query("sort")
	direction := ctx.Query("direction")

	request := struct {
		Sort      string
		Direction string
	}{
		sort,
		direction,
	}

	orders, err := eh.service.GetOrders(ctx, request)
	if err != nil {
		eh.logger.Errorf("failed to GetAllOrders err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (eh *EndpointHandler) GetOrderByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	order, err := eh.service.GetOrderByCode(ctx, code)
	if err != nil {
		eh.logger.Errorf("failed to GetOrderByCode err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (eh *EndpointHandler) ReceiveOrder(ctx *gin.Context) {
	code := ctx.Param("code")

	request := struct {
		WarehouseID int `json:"warehouse_id"`
		Items       []struct {
			ProductID int `json:"product_id"`
		} `json:"items"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		eh.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	place, err := eh.service.ReceiveOrder(ctx, pickup.ReceiveOrderRequest{
		WarehouseID: request.WarehouseID,
		OrderCode:   code,
		Items:       request.Items,
	})
	if err != nil {
		eh.logger.Errorf("failed to ReceiveOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	response := struct {
		PlaceNum int `json:"place_num"`
	}{
		place,
	}

	ctx.JSON(http.StatusOK, response)
}
