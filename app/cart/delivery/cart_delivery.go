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
	router := rg.Group("/cart")

	router.GET("/", delivery.GetUserCart)
	router.PUT("/", delivery.UpdateUserCart)
	router.DELETE("/:product_id", delivery.DeleteCartItem)
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

func (d *cartDelivery) UpdateUserCart(c *gin.Context) {
	token, err := jwt.ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}
	claim, err := jwt.ParseJWTClaims(token, false)
	userID := claim["sub"]
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	var payload domain.CartPayload
	if c.ShouldBind(&payload) != nil || payload.ProductID == "" || payload.Quantity == 0 {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, "invalid param"))
		return
	}

	err = d.handler.UpdateUserCart(c, userID.(string), payload)
	if err != nil {
		if err.Error() == "cart not found" {
			c.JSON(http.StatusNotFound, utils.HttpResponseError(http.StatusNotFound, err.Error()))
		} else {
			c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil))
}

func (d *cartDelivery) DeleteCartItem(c *gin.Context) {
	token, err := jwt.ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}
	claim, err := jwt.ParseJWTClaims(token, false)
	userID := claim["sub"]
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	productID := c.Param("product_id")
	err = d.handler.DeleteCartItem(c, userID.(string), productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil))
}
