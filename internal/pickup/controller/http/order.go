package http

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (eh *EndpointHandler) CreateOrder(ctx *gin.Context) {
	request := pickup.CreateOrderRequest{}

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

func (eh *EndpointHandler) DeleteOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	orderCode, err := eh.service.DeleteOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to DeleteOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("order with code %s: deleted", orderCode))
}

func (eh *EndpointHandler) GetOrders(ctx *gin.Context) {
	sort := ctx.Query("sort")
	direction := ctx.Query("direction")

	query := struct {
		Sort      string
		Direction string
	}{
		sort,
		direction,
	}

	orders, err := eh.service.GetOrders(ctx, query)
	if err != nil {
		eh.logger.Errorf("failed to GetAllOrders err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (eh *EndpointHandler) GetOrderByCode(ctx *gin.Context) {
	code := ctx.Param("order_code")

	order, err := eh.service.GetOrderByCode(ctx, code)
	if err != nil {
		eh.logger.Errorf("failed to GetOrderByCode err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (eh *EndpointHandler) PickupOrder(ctx *gin.Context) {
	code := ctx.Param("order_code")

	err := eh.service.PickupOrder(ctx, code)
	if err != nil {
		eh.logger.Errorf("failed to Pickup err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (eh *EndpointHandler) ReceiveOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := eh.service.ReceiveOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to ReceiveOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (eh *EndpointHandler) RefundOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := eh.service.RefundOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to RefundOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (eh *EndpointHandler) CancelOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := eh.service.CancelOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to CancelOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "canceled")
}
