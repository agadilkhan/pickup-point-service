package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (eh *EndpointHandler) initCompanyRoutes(api *gin.RouterGroup) {
	companies := api.Group("/companies")
	{
		companies.GET("/", eh.GetAllCompanies)
		companies.GET("/:company_id", eh.GetCompanyByID)
	}
}

func (eh *EndpointHandler) GetCompanyByID(ctx *gin.Context) {
	val := ctx.Param("company_id")

	id, err := strconv.Atoi(val)
	if err != nil {
		eh.logger.Errorf("failed to convert to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	company, err := eh.service.GetCompanyByID(ctx, id)
	if err != nil {
		eh.logger.Errorf("failed to GetCompanyByID err: %v", err)
		ctx.Status(http.StatusBadRequest)

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
