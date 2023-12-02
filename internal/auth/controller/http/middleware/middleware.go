package middleware

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func AuthMiddleware(cfg *config.Config, logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			logger.Errorf("missing key 'Authorization' on header")
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}
		tokenFields := strings.Fields(tokenHeader)
		if len(tokenFields) == 2 && tokenFields[0] == "Bearer" {
			tokenString = tokenFields[1]
		} else {
			logger.Errorf("incorrect token format")
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
			logger.Errorf("invalid token")
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		roleID, ok := claims["role_id"].(float64)
		if !ok {
			logger.Errorf("role_id could not parsed from jwt")
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		if roleID != 1 {
			logger.Errorf("the user does not have acces to resources")
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			logger.Errorf("user_id could not parsed from jwt")
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.Set("user_id", userID)

		ctx.Next()
	}
}
