package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func (eh *EndpointHandler) GetAllCompanies(ctx *gin.Context) {
	companies, err := eh.service.GetAllCompanies(ctx)
	if err != nil {
		eh.logger.Errorf("failed to GetAllCompanies err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, companies)
}
