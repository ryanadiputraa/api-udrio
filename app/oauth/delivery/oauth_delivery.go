package delivery

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/oauth"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
	"github.com/spf13/viper"
)

type oAuthDevlivery struct {
	handler domain.IOAuthHandler
}

func NewOAuthDelivery(rg *gin.RouterGroup, handler domain.IOAuthHandler) {
	delivery := &oAuthDevlivery{
		handler: handler,
	}
	rg.GET("/login/google", delivery.LoginGoogle)
	rg.GET("/callback", delivery.Callback)
}

func (d *oAuthDevlivery) LoginGoogle(c *gin.Context) {
	url := oauth.GetGoogleOauthConfig().AuthCodeURL(viper.GetString("OAUTH_STATE"))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (d *oAuthDevlivery) Callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if state != viper.GetString("OAUTH_STATE") {
		log.Error("state is not valid")
		oauth.RedirectWithError(c, "state is not valid")
		return
	}

	user, err := d.handler.HandleCallback(c, code)
	if err != nil {
		oauth.RedirectWithError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, user))
}
