package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (eh *EndpointHandler) JWTVerify(passwordSecretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string

		tokenHeader := ctx.Request.Header.Get("Authorization")

		tokenFields := strings.Fields(tokenHeader)
		if len(tokenFields) == 2 && tokenFields[0] == "Bearer" {
			tokenString = tokenFields[1]
		} else {
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		claims, err := eh.authService.ValidateToken(tokenString)
		if err != nil {
			eh.logger.Errorf("ValidateToken err: %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"]

		if !ok {
			eh.logger.Errorf("user_id could not parse from JWT")
		}

		ctx.Set("user_id", userID)

		ctx.Next()
	}
}
