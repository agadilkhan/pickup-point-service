package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
func (h *EndpointHandler) PickupOrder(ctx *gin.Context) {
	code := ctx.Param("order_code")

	err := h.service.PickupOrder(ctx, code)
	if err != nil {
		h.logger.Errorf("failed to Pickup err: %v", err)
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
func (h *EndpointHandler) ReceiveOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := h.service.ReceiveOrder(ctx, param)
	if err != nil {
		h.logger.Errorf("failed to ReceiveOrder err: %v", err)
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
func (h *EndpointHandler) RefundOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := h.service.RefundOrder(ctx, param)
	if err != nil {
		h.logger.Errorf("failed to RefundOrder err: %v", err)
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
func (h *EndpointHandler) CancelOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	err := h.service.CancelOrder(ctx, param)
	if err != nil {
		h.logger.Errorf("failed to CancelOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: cancelled", param),
	})
}
