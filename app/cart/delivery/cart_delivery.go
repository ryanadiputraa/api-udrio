package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type cartDelivery struct {
	handler domain.ICartHandler
}

func NewCartDelivery(rg *gin.RouterGroup, handler domain.ICartHandler) {
	delivery := &cartDelivery{
		handler: handler,
	}

	rg.GET("/cart", delivery.GetUserCart)
}

func (d *cartDelivery) GetUserCart(c *gin.Context) {
	token, err := jwt.ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	claim, err := jwt.ParseJWTClaims(token, false)
	userID := claim["sub"]
	if err != nil || userID == nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	cart, err := d.handler.GetUserCart(c, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, cart))
}
