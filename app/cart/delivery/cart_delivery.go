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
	router := rg.Group("/carts")

	router.GET("/", delivery.GetUserCart)
	router.PUT("/", delivery.UpdateUserCart)
	router.DELETE("/:product_id", delivery.DeleteCartItem)
}

func (d *cartDelivery) GetUserCart(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusForbidden, utils.HttpResponseError(http.StatusForbidden, err.Error()))
		return
	}

	cart, err := d.handler.GetUserCart(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, cart))
}

func (d *cartDelivery) UpdateUserCart(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	var payload domain.CartPayload
	if c.ShouldBind(&payload) != nil || payload.ProductID == "" || payload.Quantity == 0 {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, "invalid param"))
		return
	}

	err = d.handler.UpdateUserCart(c, userID, payload)
	if err != nil {
		if err.Error() == "cart not found" {
			c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil))
}

func (d *cartDelivery) DeleteCartItem(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	productID := c.Param("product_id")
	err = d.handler.DeleteCartItem(c, userID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil))
}
