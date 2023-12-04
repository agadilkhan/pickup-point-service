package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
)

// swagger:route POST /v1/orders/{order_code}/items/{product_id}/refund RefundItem
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
//				+ name: product_id
//	         in: path
//				+ name: Refund item request
//				in: body
//	         type: RefundItemRequest
//
//				Security:
//				  Bearer:
//
//			Responses:
//		 200: ResponseMessage
//		 400:
//		 500:
func (h *EndpointHandler) RefundItem(ctx *gin.Context) {
	params := ctx.Params

	orderCode, ok := params.Get("order_code")
	if !ok {
		h.logger.Errorf("failed to parse order_code from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	productID, ok := params.Get("product_id")
	if !ok {
		h.logger.Errorf("failed to parse product_id from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	pID, err := strconv.Atoi(productID)
	if err != nil {
		h.logger.Errorf("failed to convert product_id to int")
		ctx.Status(http.StatusBadRequest)

		return
	}

	request := struct {
		Quantity int `json:"quantity"`
	}{}

	if err = ctx.ShouldBindJSON(&request); err != nil {
		h.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = h.service.RefundItem(ctx, orderCode, pickup.RefundItemRequest{
		ProductID: pID,
		Quantity:  request.Quantity,
	})
	if err != nil {
		h.logger.Errorf("failed to RefundItem err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("%d item with product_id %s: returned", request.Quantity, productID),
	})
}

// swagger:route POST /v1/orders/{order_code}/items/{product_id}/receive ReceiveItem
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
//			+ name: product_id
//			in: path
//
//			Security:
//			  Bearer:
//
//		Responses:
//	 200: ResponseMessage
//	 400:
//	 500:
func (h *EndpointHandler) ReceiveItem(ctx *gin.Context) {
	params := ctx.Params

	orderCode, ok := params.Get("order_code")
	if !ok {
		h.logger.Errorf("failed to parse order_code from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	productID, ok := params.Get("product_id")
	if !ok {
		h.logger.Errorf("failed to parse product_id from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	pID, err := strconv.Atoi(productID)
	if err != nil {
		h.logger.Errorf("failed to convert product_id to int")
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = h.service.ReceiveItem(ctx, orderCode, pID)
	if err != nil {
		h.logger.Errorf("failed to ReceiveItem err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("item with product_id %s: received", productID),
	})
}
