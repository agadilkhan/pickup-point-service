package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (eh *EndpointHandler) initPickupRoutes(api *gin.RouterGroup) {
	pickupOrders := api.Group("/:user_id/pickup_orders")
	{
		pickupOrders.GET("/", eh.GetPickupOrders)
		pickupOrders.GET("/:pickup_order_id", eh.GetPickupOrderByID)
	}
	pickupPoints := api.Group("/pickup_points")
	{
		pickupPoints.GET("/", eh.GetAllPickupPoints)
		pickupPoints.GET("/:pickup_point_id", eh.GetPickupPointByID)
	}
}

func (eh *EndpointHandler) GetPickupOrderByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		eh.logger.Errorf("failed to convert request user_id to int err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = middleware.CheckUser(ctx, userID)
	if err != nil {
		eh.logger.Errorf("failed to CheckUser err: %v", err)
		ctx.Status(http.StatusNotFound)

		return
	}

	pVal := ctx.Param("pickup_order_id")

	pickupOrderID, err := strconv.Atoi(pVal)
	if err != nil {
		eh.logger.Errorf("failed to convert pickup_order_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	request := struct {
		UserID        int
		PickupOrderID int
	}{
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

func (eh *EndpointHandler) GetPickupOrders(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		eh.logger.Errorf("failed to convert request user_id to int err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = middleware.CheckUser(ctx, userID)
	if err != nil {
		eh.logger.Errorf("failed to CheckUser err: %v", err)
		ctx.Status(http.StatusNotFound)

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

func (eh *EndpointHandler) GetAllPickupPoints(ctx *gin.Context) {
	points, err := eh.service.GetAllPickupPoints(ctx)
	if err != nil {
		eh.logger.Errorf("failed to GetAllPickupPoints err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, points)
}

func (eh *EndpointHandler) GetPickupPointByID(ctx *gin.Context) {
	val := ctx.Param("pickup_point_id")

	id, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("failed to convert pickup_point_id to int")
		ctx.Status(http.StatusBadRequest)

		return
	}

	point, err := eh.service.GetPickupPointByID(ctx, id)
	if err != nil {
		eh.logger.Errorf("failed to GetPickupPointByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, point)
}
