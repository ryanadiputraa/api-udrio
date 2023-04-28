package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type orderDelivery struct {
	handler domain.IOrderHandler
}

func NewOrderDelivery(rg *gin.RouterGroup, handler domain.IOrderHandler) {
	delivery := orderDelivery{handler: handler}
	router := rg.Group("/orders")

	router.GET("/", delivery.GetUserOrders)
	router.POST("/", delivery.CreateOrder)
}

func (d *orderDelivery) GetUserOrders(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusForbidden, utils.HttpResponseError(http.StatusForbidden, err.Error()))
		return
	}

	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	orders, meta, err := d.handler.GetUserOrders(c, userID, size, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponseWithMetaData(http.StatusOK, orders, meta))
}

func (d *orderDelivery) CreateOrder(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusForbidden, utils.HttpResponseError(http.StatusForbidden, err.Error()))
		return
	}

	var payload domain.OrderPayload
	if c.ShouldBind(&payload) != nil || len(payload.Orders) < 1 {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, "invalid param"))
		return
	}

	err = d.handler.CreateOrder(c, userID, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.HttpResponse(http.StatusCreated, nil))
}
