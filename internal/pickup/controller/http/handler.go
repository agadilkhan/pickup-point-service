package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type EndpointHandler struct {
	service pickup.UseCase
	logger  *zap.SugaredLogger
	cfg     *config.Config
}

func NewEndpointHandler(
	service pickup.UseCase,
	logger *zap.SugaredLogger,
	cfg *config.Config,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
		cfg:     cfg,
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

func (eh *EndpointHandler) Receive(ctx *gin.Context) {

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

func (eh *EndpointHandler) CreateOrder(ctx *gin.Context) {

}

func (eh *EndpointHandler) GetCustomerByID(ctx *gin.Context) {
	val := ctx.Param("id")

	id, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("cannot convert to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	customer, err := eh.service.GetCustomerByID(ctx, id)
	if err != nil {
		eh.logger.Errorf("failed to GetCustomerByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (eh *EndpointHandler) GetCompanyByID(ctx *gin.Context) {
	val := ctx.Param("id")

	id, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("cannot convert to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	company, err := eh.service.GetCompanyByID(ctx, id)
	if err != nil {
		eh.logger.Errorf("failed to GetCompanyByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, company)
}

func (eh *EndpointHandler) GetPickupOrders(ctx *gin.Context) {
	val := ctx.Param("user_id")

	userID, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("cannot convert user_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	pickupOrders, err := eh.service.GetPickupOrders(ctx, userID)
	if err != nil {
		eh.logger.Errorf("failed to GetPickupOrders err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, pickupOrders)
}

func (eh *EndpointHandler) GetPickupOrderByID(ctx *gin.Context) {
	uVal := ctx.Param("user_id")

	userID, err := strconv.Atoi(uVal)
	if err != nil {
		eh.logger.Errorf("cannot convert user_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	pVal := ctx.Param("pickup_order_id")

	pickupOrderID, err := strconv.Atoi(pVal)
	if err != nil {
		eh.logger.Errorf("cannot convert pickup_order_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	request := pickup.GetPickupOrderByIDRequest{
		userID,
		pickupOrderID,
	}

	pickupOrder, err := eh.service.GetPickupOrderByID(ctx, request)
	if err != nil {
		eh.logger.Errorf("failed to GetPickupOrderByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, pickupOrder)
}

func (eh *EndpointHandler) GetOrders(ctx *gin.Context) {
	sort := ctx.Query("sort")
	direction := ctx.Query("direction")

	request := pickup.GetAllOrdersRequest{
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
