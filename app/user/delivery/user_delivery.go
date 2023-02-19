package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
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
	user, err := d.handler.GetUserInfo(c, "")
	if err != nil {
		errMsg := map[string]string{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, utils.HttpResponse(http.StatusInternalServerError, errMsg, nil))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, user))
}
