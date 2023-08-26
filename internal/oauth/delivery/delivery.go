package delivery

import (
	"net/http"

	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/oauth"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type delivery struct {
	config      config.Config
	conf        config.Config
	usecase     domain.OAuthUsecase
	userUsecase domain.UserUsecase
}

func NewOAuthDelivery(rg *gin.RouterGroup, conf config.Config, usecase domain.OAuthUsecase, userUsecase domain.UserUsecase) {
	delivery := &delivery{
		conf:        conf,
		usecase:     usecase,
		userUsecase: userUsecase,
	}
	rg.GET("/login/google", delivery.LoginGoogle)
	rg.GET("/callback", delivery.Callback)
	rg.POST("/refresh", delivery.RefreshToken)
}

func (d *delivery) LoginGoogle(c *gin.Context) {
	url := oauth.GetGoogleOauthConfig(d.conf.Oauth).AuthCodeURL(d.conf.Oauth.State, oauth2.SetAuthURLParam("prompt", "select_account"))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (d *delivery) Callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if state != d.conf.Oauth.State {
		oauth.RedirectWithError(c, d.conf.Oauth.RedirectURL, "state is not valid")
		return
	}

	user, err := d.usecase.HandleCallback(c, d.conf.Oauth, code)
	if err != nil {
		oauth.RedirectWithError(c, d.conf.Oauth.RedirectURL, err.Error())
		return
	}

	// save or update user
	userData := domain.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Picture:   user.Picture,
		Locale:    user.Locale,
	}

	err = d.userUsecase.CreateOrUpdateIfExist(c, userData)
	if err != nil {
		oauth.RedirectWithError(c, d.conf.Oauth.RedirectURL, err.Error())
		return
	}

	// generate access token
	tokens, err := d.usecase.GenerateAccessToken(c, userData.ID)
	if err != nil {
		oauth.RedirectWithError(c, d.conf.Oauth.RedirectURL, err.Error())
		return
	}

	oauth.RedirectWithAccessToken(c, d.conf.Oauth.RedirectURL, tokens)
}

func (d *delivery) RefreshToken(c *gin.Context) {
	refreshToken, err := jwt.ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}

	tokens, err := d.usecase.RefreshAccessToken(c, refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.HttpResponseError(http.StatusUnauthorized, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, tokens))
}
