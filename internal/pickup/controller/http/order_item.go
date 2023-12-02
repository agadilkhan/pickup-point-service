package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (eh *EndpointHandler) RefundItem(ctx *gin.Context) {
	params := ctx.Params

	orderCode, ok := params.Get("order_code")
	if !ok {
		eh.logger.Errorf("failed to parse order_code from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	productID, ok := params.Get("product_id")
	if !ok {
		eh.logger.Errorf("failed to parse product_id from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	pID, err := strconv.Atoi(productID)
	if err != nil {
		eh.logger.Errorf("failed to convert product_id to int")
		ctx.Status(http.StatusBadRequest)

		return
	}

	request := struct {
		Quantity int `json:"quantity"`
	}{}

	if err = ctx.ShouldBindJSON(&request); err != nil {
		eh.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = eh.service.RefundItem(ctx, orderCode, pickup.RefundItemRequest{
		ProductID: pID,
		Quantity:  request.Quantity,
	})
	if err != nil {
		eh.logger.Errorf("failed to RefundItem err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (eh *EndpointHandler) ReceiveItem(ctx *gin.Context) {
	params := ctx.Params

	orderCode, ok := params.Get("order_code")
	if !ok {
		eh.logger.Errorf("failed to parse order_code from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	productID, ok := params.Get("product_id")
	if !ok {
		eh.logger.Errorf("failed to parse product_id from params")
		ctx.Status(http.StatusBadRequest)

		return
	}

	pID, err := strconv.Atoi(productID)
	if err != nil {
		eh.logger.Errorf("failed to convert product_id to int")
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = eh.service.ReceiveItem(ctx, orderCode, pID)
	if err != nil {
		eh.logger.Errorf("failed to ReceiveItem err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}
