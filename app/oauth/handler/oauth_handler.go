package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/oauth"
)

type oAuthHandler struct{}

func NewOAuthHandler() domain.IOAuthHandler {
	return &oAuthHandler{}
}

func (h *oAuthHandler) HandleCallback(ctx context.Context, code string) (user domain.GoogleProfile, err error) {
	token, err := oauth.GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		log.Error("failed to retrieve token", err.Error())
		return user, errors.New("failed to retrive token")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Error("failed to retrive user info", err.Error())
		return user, errors.New("failed to retrive user info")
	}
	defer resp.Body.Close()

	var userInfo domain.GoogleProfile
	json.NewDecoder(resp.Body).Decode(&userInfo)

	return userInfo, nil
}

func (h *oAuthHandler) GenerateAccessToken(ctx context.Context, userID interface{}) (domain.Tokens, error) {
	return domain.Tokens{
		AccessToken:  "access_token",
		ExpiresIn:    3600,
		RefreshToken: "refresh_token",
	}, nil
}
