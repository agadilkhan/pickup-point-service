package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (eh *EndpointHandler) initProductRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", eh.GetAllProducts)
		products.GET("/:product_id", eh.GetProductByID)
	}
}

func (eh *EndpointHandler) GetAllProducts(ctx *gin.Context) {
	products, err := eh.service.GetAllProducts(ctx)
	if err != nil {
		eh.logger.Errorf("failed to GetAllProducts err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (eh *EndpointHandler) GetProductByID(ctx *gin.Context) {
	val := ctx.Param("product_id")

	id, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("failed to convert product_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	product, err := eh.service.GetProductByID(ctx, id)
	if err != nil {
		eh.logger.Errorf("failed to GetProductByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, product)
}
