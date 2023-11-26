package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func (eh *EndpointHandler) GetPickupOrderByID(ctx *gin.Context) {
	userID, err := middleware.CheckUser(ctx)
	if err != nil {
		eh.logger.Errorf("failed to CheckUser err: %v", err)
		ctx.Status(http.StatusBadRequest)

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
	userID, err := middleware.CheckUser(ctx)
	if err != nil {
		eh.logger.Errorf("failed to CheckUser err: %v", err)
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
