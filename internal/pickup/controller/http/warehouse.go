package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (eh *EndpointHandler) initWarehouseRoutes(api *gin.RouterGroup) {
	warehouses := api.Group("/warehouses")
	{
		warehouses.GET("/", eh.GetWarehouses)
		warehouses.GET("/:warehouse_id", eh.GetWarehouseByID)
	}
}

func (eh *EndpointHandler) GetWarehouses(ctx *gin.Context) {
	warehouses, err := eh.service.GetWarehouses(ctx)
	if err != nil {
		eh.logger.Errorf("failed to GetAllWarehouses err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, warehouses)
}

func (eh *EndpointHandler) GetWarehouseByID(ctx *gin.Context) {
	val := ctx.Param("warehouse_id")

	id, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("failed to convert warehouse_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	warehouse, err := eh.service.GetWarehouseByID(ctx, id)
	if err != nil {
		eh.logger.Errorf("failed to GetWarehouseByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, warehouse)
}
