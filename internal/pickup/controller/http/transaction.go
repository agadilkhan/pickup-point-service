package http

import (
	"net/http"
	"strconv"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /v1/{user_id}/transactions Transactions
//
//				Consumes:
//				- application/json
//
//				Produces:
//				- application/json
//
//				Schemes: http, https
//
//				Parameters:
//					+ name: user_id
//					in: path
//					+ name: transaction_type
//		         in: query
//
//					Security:
//					  Bearer:
//
//				Responses:
//			 200: ResponseOK
//			 400:
//	      404:
//			 500:
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
		TransactionType: query,
	}

	transactions, err := eh.service.GetTransactions(ctx, userID, transactionQuery)
	if err != nil {
		eh.logger.Errorf("failed to GetTransaction err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseOK{
		Data: transactions,
	})
}
