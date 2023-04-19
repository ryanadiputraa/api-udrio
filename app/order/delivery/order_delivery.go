package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type orderDelivery struct {
	handler domain.IOrderHandler
}

func NewOrderDelivery(rg *gin.RouterGroup, handler domain.IOrderHandler) {
	deliver := orderDelivery{handler: handler}
	router := rg.Group("/orders")

	router.GET("/", deliver.GetUserOrders)
}

func (o *orderDelivery) GetUserOrders(c *gin.Context) {
	token, err := jwt.ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	claim, err := jwt.ParseJWTClaims(token, false)
	userID := claim["sub"]
	if err != nil || userID == nil {
		c.JSON(http.StatusForbidden, utils.HttpResponseError(http.StatusForbidden, err.Error()))
		return
	}

	orders, err := o.handler.GetUserOrders(c, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, orders))
}
