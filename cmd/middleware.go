package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtUtils "github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := jwtUtils.ExtractTokenFromAuthorizationHeader(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponse(http.StatusUnauthorized, err.Error(), nil))
			return
		}

		_, err = jwtUtils.ParseJWTClaims(tokenString, false)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponse(http.StatusUnauthorized, err.Error(), nil))
			return
		}

		ctx.Next()
	}
}
