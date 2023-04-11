package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type userDelivery struct {
	handler domain.IUserHandler
}

func NewUserDelivery(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc, handler domain.IUserHandler) {
	delivery := userDelivery{handler: handler}
	router := rg.Group("/users")

	router.GET("/", authMiddleware, delivery.GetUserInfo)
}

func (d *userDelivery) GetUserInfo(c *gin.Context) {
	token, err := jwt.ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}

	claims, err := jwt.ParseJWTClaims(token, false)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}

	user, err := d.handler.GetUserInfo(c, claims["sub"])
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, user))
}
