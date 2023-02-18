package domain

import (
	"context"
)

type IOAuthHandler interface {
	HandleCallback(ctx context.Context, code string) (user GoogleProfile, err error)
	GenerateAccessToken(ctx context.Context, userID interface{}) (Tokens, error)
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type GoogleProfile struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Locale  string `json:"locale"`
}
