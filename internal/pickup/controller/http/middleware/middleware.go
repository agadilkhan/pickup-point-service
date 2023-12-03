package middleware

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/metrics"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTVerify(cfg *config.Config, logger *zap.SugaredLogger) gin.HandlerFunc {
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

		userID, ok := claims["user_id"]
		if !ok {
			logger.Errorf("user_id could not parsed from jwt")
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.Set("user_id", userID)

		ctx.Next()
	}
}

func CheckUser(ctx *gin.Context, requestUser int) error {
	contextUser, ok := ctx.Value("user_id").(float64)
	if !ok {
		return fmt.Errorf("failed to convert context user_id to float64")
	}

	if int(contextUser) != requestUser {
		return fmt.Errorf("the user does not have access to this resource")
	}

	return nil
}

func MetricsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		path := ctx.Request.URL.Path

		statusString := strconv.Itoa(ctx.Writer.Status())

		metrics.HttpResponseTime.WithLabelValues(path, statusString, ctx.Request.Method).Observe(time.Since(start).Seconds())
		metrics.HttpRequestTotalCollector.WithLabelValues(path, statusString, ctx.Request.Method).Inc()
	}
}
