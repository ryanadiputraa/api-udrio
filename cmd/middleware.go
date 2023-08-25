package main

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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
			return
		}

		_, err = jwtUtils.ParseJWTClaims(tokenString, false)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.Next()
	}
}
