package http

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/pkg/pagination"
	"net/http"
	"strconv"
	"strings"
	"time"

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
func (h *EndpointHandler) CreateOrder(ctx *gin.Context) {
	request := pickup.CreateOrderRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	orderID, err := h.service.CreateOrder(ctx, request)
	if err != nil {
		h.logger.Errorf("failed to CreateOrder err: %v", err)
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
func (h *EndpointHandler) DeleteOrder(ctx *gin.Context) {
	param := ctx.Param("order_code")

	orderCode, err := h.service.DeleteOrder(ctx, param)
	if err != nil {
		h.logger.Errorf("failed to DeleteOrder err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("order with code %s: deleted", orderCode),
	})
}

func (h *EndpointHandler) GetOrders(ctx *gin.Context) {
	sortBy := ctx.Query("sort_by")
	sortOrder := ctx.Query("sort_order")

	sortOptions := pagination.SortOptions{
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	var filterOptions pagination.FilterOptions

	var totalAmountBuilder = pagination.NewFieldBuilder(&pagination.Field{})
	totalAmount := ctx.Query("total_amount")
	if totalAmount != "" {
		value := totalAmount
		totalAmountBuilder.SetName("total_amount")
		if strings.Index(value, ":") != -1 {
			split := strings.Split(totalAmount, ":")
			if _, err := strconv.ParseFloat(split[0], 64); err == nil {
				totalAmountBuilder.SetValues(split[0], split[1])
				totalAmountBuilder.SetOperator("between")
			} else {
				totalAmountBuilder.SetOperator(split[0])
				totalAmountBuilder.SetValues(split[1])
			}
		} else {
			totalAmountBuilder.SetOperator("eq")
			totalAmountBuilder.SetValues(value)
		}
		totalAmountField := totalAmountBuilder.Build()
		filterOptions.AddField(*totalAmountField)
	}

	var createdAtBuilder = pagination.NewFieldBuilder(&pagination.Field{})
	createdAt := ctx.Query("created_at")
	if createdAt != "" {
		value := createdAt
		createdAtBuilder.SetName("created_at")
		if strings.Index(value, ":") != -1 {
			split := strings.Split(createdAt, ":")
			startDate, err := time.Parse("2006-01-02", split[0])
			if err != nil {
				h.logger.Errorf("incorrect date format")
				ctx.Status(http.StatusBadRequest)

				return
			}
			endDate, err := time.Parse("2006-01-02", split[1])
			if err != nil {
				h.logger.Errorf("incorrect date format")
				ctx.Status(http.StatusBadRequest)

				return
			}
			createdAtBuilder.SetValues(startDate, endDate)
			createdAtBuilder.SetOperator("between")
		} else {
			date, err := time.Parse("2006-01-02", value)
			if err != nil {
				h.logger.Errorf("incorrect date format")
				ctx.Status(http.StatusBadRequest)

				return
			}
			createdAtBuilder.SetValues(date)
			createdAtBuilder.SetOperator("eq")
		}
		createdAtField := createdAtBuilder.Build()
		filterOptions.AddField(*createdAtField)
	}

	orders, err := h.service.GetOrders(ctx, sortOptions, filterOptions)
	if err != nil {
		h.logger.Errorf("failed to GetAllOrders err: %v", err)
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
func (h *EndpointHandler) GetOrderByCode(ctx *gin.Context) {
	code := ctx.Param("order_code")

	order, err := h.service.GetOrderByCode(ctx, code)
	if err != nil {
		h.logger.Errorf("failed to GetOrderByCode err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseOK{
		Data: order,
	})
}
