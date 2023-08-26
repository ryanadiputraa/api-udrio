package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/config"
	jwtUtils "github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

func AuthMiddleware(conf config.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := jwtUtils.ExtractTokenFromAuthorizationHeader(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
			return
		}

		_, err = jwtUtils.ParseJWTClaims(conf, tokenString, false)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.Next()
	}
}
