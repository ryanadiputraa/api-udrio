package domain

import (
	"context"
)

type IOAuthService interface {
	HandleCallback(ctx context.Context, code string) (user GoogleProfile, err error)
}

type GoogleProfile struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Locale  string `json:"locale"`
}
