package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type delivery struct {
	usecase domain.CartUsecase
}

func NewCartDelivery(rg *gin.RouterGroup, usecase domain.CartUsecase) {
	delivery := &delivery{
		usecase: usecase,
	}
	router := rg.Group("/carts")

	router.GET("/", delivery.GetUserCart)
	router.PUT("/", delivery.UpdateUserCart)
	router.DELETE("/:product_id", delivery.DeleteCartItem)
}

func (d *delivery) GetUserCart(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusForbidden, utils.HttpResponseError(http.StatusForbidden, err.Error()))
		return
	}

	cart, err := d.usecase.GetUserCart(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, cart))
}

func (d *delivery) UpdateUserCart(c *gin.Context) {
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

	err = d.usecase.UpdateUserCart(c, userID, payload)
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

func (d *delivery) DeleteCartItem(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	productID := c.Param("product_id")
	err = d.usecase.DeleteCartItem(c, userID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil))
}
