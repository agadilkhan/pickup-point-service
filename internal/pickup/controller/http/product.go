package http

import (
	"net/http"

	"github.com/agadilkhan/pickup-point-service/pkg/pagination"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /v1/products GetProducts
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
//				+ name: name
//	         in: query
//
//				Security:
//				  Bearer:
//
//			Responses:
//		 200: ResponseOK
//		 500:
func (h *EndpointHandler) GetProducts(ctx *gin.Context) {
	name := ctx.Query("name")

	searchOptions := pagination.SearchOptions{
		Field: "name",
		Value: name,
	}

	products, err := h.service.GetProducts(ctx, searchOptions)
	if err != nil {
		h.logger.Errorf("failed to GetProducst err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseOK{
		Data: products,
	})
}
