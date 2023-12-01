package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (eh *EndpointHandler) initPickupPointRoutes(api *gin.RouterGroup) {
	pickupPoints := api.Group("/pickup_points")
	{
		pickupPoints.GET("/", eh.GetPickupPoints)
		pickupPoints.GET("/:pickup_point_id", eh.GetPickupPointByID)
	}
}

func (eh *EndpointHandler) GetPickupPoints(ctx *gin.Context) {
	points, err := eh.service.GetPickupPoints(ctx)
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
