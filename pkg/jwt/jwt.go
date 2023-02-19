package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ExtractTokenFromAuthorizationHeader(c *gin.Context) (token string, err error) {
	t := c.GetHeader("Authorization")
	if len(t) == 0 {
		return "", errors.New("missing authorization header")
	}

	h := strings.Split(t, " ")
	if len(h) < 2 || h[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	token = h[1]
	return token, nil
}
