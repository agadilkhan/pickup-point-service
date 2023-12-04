package http

import (
	"fmt"
	"net/http"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
)

// swagger:route POST /v1/orders/ CreateOrder
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//		+ name: CreateOrderRequest
//			in: body
//			type: CreateOrderRequest
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 201: ResponseCreated
//	 400:
//	 500:
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

	ctx.JSON(http.StatusCreated, responseCreated{
		ID: orderID,
	})
}

// swagger:route DELETE /v1/orders/{order_code} DeleteOrder
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//			+ name: order_code
//			in: path
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 200: ResponseMessage
//	 500:
func (eh *EndpointHandler) DeleteOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	orderCode, err := eh.service.DeleteOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to DeleteOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: deleted", orderCode),
	})
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

// swagger:route GET /v1/orders/{order_code} GetOrderByCode
//
//			Consumes:
//			- application/json
//
//			Produces:
//			- application/json
//
//			Schemes: http, https
//
//			Parameters:
//				+ name: order_code
//				in: path
//
//			Security:
//			  Bearer:
//
//			Responses:
//		 200: ResponseOK
//	  500:
func (eh *EndpointHandler) GetOrderByCode(ctx *gin.Context) {
	code := ctx.Param("order_code")

	order, err := eh.service.GetOrderByCode(ctx, code)
	if err != nil {
		eh.logger.Errorf("failed to GetOrderByCode err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseOK{
		Data: order,
	})
}

// swagger:route POST /v1/orders/{order_code}/pickup PickupOrder
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//			+ name: order_code
//			in: path
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 200: ResponseMessage
//	 400:
//	 500:
func (eh *EndpointHandler) PickupOrder(ctx *gin.Context) {
	code := ctx.Param("order_code")

	err := eh.service.PickupOrder(ctx, code)
	if err != nil {
		eh.logger.Errorf("failed to Pickup err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: issued", code),
	})
}

// swagger:route POST /v1/orders/{order_code}/receive ReceiveOrder
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//			+ name: order_code
//			in: path
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 200: ResponseMessage
//	 400:
//	 500:
func (eh *EndpointHandler) ReceiveOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := eh.service.ReceiveOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to ReceiveOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: received", param),
	})
}

// swagger:route POST /v1/orders/{order_code}/refund RefundOrder
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//			+ name: order_code
//			in: path
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 200: ResponseMessage
//	 400:
//	 500:
func (eh *EndpointHandler) RefundOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := eh.service.RefundOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to RefundOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: returned", param),
	})
}

// swagger:route PUT /v1/orders/{order_code}/cancel CancelOrder
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//			+ name: order_code
//			in: path
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 200: ResponseMessage
//	 400:
//	 500:
func (eh *EndpointHandler) CancelOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := eh.service.CancelOrder(ctx, param)
	if err != nil {
		eh.logger.Errorf("failed to CancelOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: cancelled", param),
	})
}
