package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTVerify(cfg *config.Config) gin.HandlerFunc {
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

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to SigningMethodHMAC")
			}

			return []byte(cfg.Auth.JWTSecretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			log.Printf("user_id could not parsed from JWT")

			return
		}

		ctx.Set("user_id", userID)

		ctx.Next()
	}
}
