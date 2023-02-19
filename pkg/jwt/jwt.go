package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ryanadiputraa/api-udrio/domain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GenerateAccessToken(userID interface{}) (tokens domain.Tokens, err error) {
	// access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(viper.GetDuration("JWT_EXPIRES_IN")).Unix(),
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		log.Error("failed to generate access token:", err.Error())
		return tokens, err
	}

	// refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(viper.GetDuration("JWT_REFRESH_EXPIRES_IN")).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(viper.GetString("JWT_REFRESH_SECRET")))
	if err != nil {
		log.Error("failed to generate refresh token:", err.Error())
		return tokens, err
	}

	tokens = domain.Tokens{
		AccessToken:  tokenString,
		ExpiresIn:    int(time.Now().Add(viper.GetDuration("JWT_EXPIRES_IN")).Unix()),
		RefreshToken: refreshTokenString,
	}

	return tokens, nil
}
