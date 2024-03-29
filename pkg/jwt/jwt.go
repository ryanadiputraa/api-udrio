package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
)

func GenerateAccessToken(conf config.JWT, userID interface{}) (tokens domain.Tokens, err error) {
	// access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(conf.ExpiresIn).Unix(),
	})
	tokenString, err := token.SignedString([]byte(conf.Secret))
	if err != nil {
		return tokens, err
	}

	// refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(conf.RefreshExpiresIn).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(conf.RefreshSecret))
	if err != nil {
		return tokens, err
	}

	tokens = domain.Tokens{
		AccessToken:  tokenString,
		ExpiresIn:    int(time.Now().Add(conf.ExpiresIn).Unix()),
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

func ParseJWTClaims(conf config.JWT, tokenString string, isRefresh bool) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		var secret string
		if !isRefresh {
			secret = conf.Secret
		} else {
			secret = conf.RefreshSecret
		}

		return []byte(secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ExtractUserID(c *gin.Context, conf config.JWT) (userID string, err error) {
	token, err := ExtractTokenFromAuthorizationHeader(c)
	if err != nil {
		return
	}

	claim, err := ParseJWTClaims(conf, token, false)
	userID = claim["sub"].(string)
	return
}
