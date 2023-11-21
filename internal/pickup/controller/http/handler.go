package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
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
	logger := eh.logger.With(
		zap.String("endpoint", "pickup"),
		zap.String("params", ctx.FullPath()),
	)

	code := ctx.Param("code")

	err := eh.service.Pickup(ctx, code)
	if err != nil {
		logger.Errorf("failed to Pickup err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (eh *EndpointHandler) GetOrderByCode(ctx *gin.Context) {
	logger := eh.logger.With(
		zap.String("endpoint", "get order by code"),
		zap.String("params", ctx.FullPath()),
	)

	code := ctx.Param("code")

	order, err := eh.service.GetOrderByCode(ctx, code)
	if err != nil {
		logger.Errorf("failed to GetOrderByCode err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	response := struct {
		ID            int
		CustomerID    int
		CompanyID     int
		PointID       int
		Code          string
		Status        entity.OrderStatus
		PaymentStatus entity.PaymentStatus
		Products      []entity.OrderItem
		TotalAmount   float64
	}{
		order.ID,
		order.CustomerID,
		order.CompanyID,
		order.PointID,
		order.Code,
		order.Status,
		order.PaymentStatus,
		order.OrderItems,
		order.TotalAmount,
	}

	ctx.JSON(http.StatusOK, response)
}

func (eh *EndpointHandler) CreateOrder(ctx *gin.Context) {

}

func (eh *EndpointHandler) GetCustomerByID(ctx *gin.Context) {
	logger := eh.logger.With(
		zap.String("endpoint", "get customer by id"),
		zap.String("params", ctx.FullPath()),
	)

	val := ctx.Param("id")

	id, err := strconv.Atoi(val)
	if err != nil {
		logger.Errorf("cannot convert to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	customer, err := eh.service.GetCustomerByID(ctx, id)
	if err != nil {
		logger.Errorf("failed to GetCustomerByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (eh *EndpointHandler) GetCompanyByID(ctx *gin.Context) {
	logger := eh.logger.With(
		zap.String("endpoint", "get customer by id"),
		zap.String("params", ctx.FullPath()),
	)

	val := ctx.Param("id")

	id, err := strconv.Atoi(val)
	if err != nil {
		logger.Errorf("cannot convert to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	company, err := eh.service.GetCompanyByID(ctx, id)
	if err != nil {
		logger.Errorf("failed to GetCompanyByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, company)
}
