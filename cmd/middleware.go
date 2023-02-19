package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	jwtUtils "github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info("auth middleware")

		tokenString, err := jwtUtils.ExtractTokenFromAuthorizationHeader(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponse(http.StatusUnauthorized, err.Error(), nil))
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponse(http.StatusUnauthorized, err.Error(), nil))
			return
		}

		ctx.Next()
	}
}
