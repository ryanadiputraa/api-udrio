package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type delivery struct {
	usecase domain.UserUsecase
}

func NewUserDelivery(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc, usecase domain.UserUsecase) {
	delivery := delivery{usecase: usecase}
	router := rg.Group("/users")

	router.GET("/", authMiddleware, delivery.GetUserInfo)
}

func (d *delivery) GetUserInfo(c *gin.Context) {
	userID, err := jwt.ExtractUserID(c)
	if err != nil || userID == "" {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}

	user, err := d.usecase.GetUserInfo(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, user))
}
