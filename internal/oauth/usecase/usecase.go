package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/jwt"
	"github.com/ryanadiputraa/api-udrio/pkg/oauth"
)

type usecase struct{}

func NewOAuthUsecase() domain.OAuthUsecase {
	return &usecase{}
}

func (u *usecase) HandleCallback(ctx context.Context, conf config.Oauth, code string) (user domain.GoogleProfile, err error) {
	token, err := oauth.GetGoogleOauthConfig(conf).Exchange(context.Background(), code)
	if err != nil {
		return user, errors.New("fail to retrive token")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return user, errors.New("fail to retrive user info")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *usecase) GenerateAccessToken(ctx context.Context, userID interface{}) (domain.Tokens, error) {
	tokens, err := jwt.GenerateAccessToken(userID)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

func (u *usecase) RefreshAccessToken(ctx context.Context, refreshToken string) (tokens domain.Tokens, err error) {
	claims, err := jwt.ParseJWTClaims(refreshToken, true)
	if err != nil {
		return tokens, err
	}

	tokens, err = jwt.GenerateAccessToken(claims["sub"])
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}
