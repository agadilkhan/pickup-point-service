package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (eh *EndpointHandler) GetTransactions(ctx *gin.Context) {
	param := ctx.Param("user_id")
	query := ctx.Query("transaction_type")

	userID, err := strconv.Atoi(param)
	if err != nil {
		eh.logger.Errorf("failed to convert user_id to int")
		ctx.Status(http.StatusBadRequest)

		return
	}

	err = middleware.CheckUser(ctx, userID)
	if err != nil {
		eh.logger.Errorf("the user does not have access to resource")
		ctx.Status(http.StatusNotFound)

		return
	}

	transactionQuery := pickup.GetTransactionsQuery{
		query,
	}

	transactions, err := eh.service.GetTransactions(ctx, userID, transactionQuery)
	if err != nil {
		eh.logger.Errorf("failed to GetTransaction err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
