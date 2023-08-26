package oauth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOauthConfig(conf config.Oauth) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  conf.CallbackURL,
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func RedirectWithError(c *gin.Context, redirectURL string, err string) {
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?err=%s", redirectURL, err))
}

func RedirectWithAccessToken(c *gin.Context, redirectURL string, tokens domain.Tokens) {
	url := fmt.Sprintf("%s?access_token=%s&expires_in=%s&refresh_token=%s", redirectURL, tokens.AccessToken, strconv.Itoa(tokens.ExpiresIn), tokens.RefreshToken)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
