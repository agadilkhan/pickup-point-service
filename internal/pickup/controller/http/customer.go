package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func (eh *EndpointHandler) GetAllCustomers(ctx *gin.Context) {
	customers, err := eh.service.GetAllCustomers(ctx)
	if err != nil {
		eh.logger.Errorf("failed to GetAllCustomers err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, customers)
}
