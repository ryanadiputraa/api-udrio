package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/oauth"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
	"github.com/spf13/viper"
)

type OauthHandler struct {
	service domain.IOAuthService
}

func NewOauthHandler(rg *gin.RouterGroup, service domain.IOAuthService) {
	handler := &OauthHandler{
		service: service,
	}
	rg.GET("/login/google", handler.LoginGoogle)
	rg.GET("/callback", handler.Callback)
}

func (h *OauthHandler) LoginGoogle(c *gin.Context) {
	url := oauth.GetGoogleOauthConfig().AuthCodeURL(viper.GetString("OAUTH_STATE"))
	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}

func (h *OauthHandler) Callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if state != viper.GetString("OAUTH_STATE") {
		oauth.RedirectWithError(c, "state is not valid")
		return
	}

	user, err := h.service.HandleCallback(c, code)
	if err != nil {
		oauth.RedirectWithError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, user))
}
