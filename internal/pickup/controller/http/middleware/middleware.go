package middleware

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var logger = zap.SugaredLogger{}

func JWTVerify(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		tokenHeader := ctx.Request.Header.Get("Authorization")
		tokenFields := strings.Fields(tokenHeader)
		if len(tokenFields) == 2 && tokenFields[0] == "Bearer" {
			tokenString = tokenFields[1]
		} else {
			logger.Error("incorrect token format")
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to SigningMethodHMAC")
			}

			return []byte(cfg.Auth.JWTSecretKey), nil
		})

		if err != nil {
			logger.Errorf("failed to ParseWithClaims err: %v", err)
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		if !token.Valid {
			logger.Error("invalid token")
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			logger.Error("user_id could not parsed from jwt")
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.Set("user_id", userID)

		ctx.Next()
	}
}

func CheckUser(ctx *gin.Context) (int, error) {
	accessUser, ok := ctx.Value("user_id").(float64)
	if !ok {
		return 0, fmt.Errorf("failed to convert user_id to float64")
	}

	requestUser, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return 0, fmt.Errorf("failed to convert user_id to int: %v", err)
	}

	if int(accessUser) != requestUser {
		return 0, fmt.Errorf("the user does not have access to this resource")
	}

	return requestUser, nil
}
