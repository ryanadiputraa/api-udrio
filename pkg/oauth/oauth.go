package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  viper.GetString("OAUTH_CALLBACK_URL"),
		ClientID:     viper.GetString("OAUTH_CLIENT_ID"),
		ClientSecret: viper.GetString("OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func RedirectWithError(c *gin.Context, err string) {
	redirectURL := viper.GetString("OAUTH_REDIRECT_URL")
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?err=%s", redirectURL, err))
}

func RedirectWithAccessToken(c *gin.Context, accessToken string) {
	redirectURL := viper.GetString("OAUTH_REDIRECT_URL")
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?access_token=%s", redirectURL, accessToken))
}
