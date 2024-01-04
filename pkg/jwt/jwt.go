package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

var (
	CustomSecret = []byte("夏天夏天悄悄过去")
)

type CustomClaims struct {
	Username string `json:"username"`
	UserId   int64  `json:"user_id"`
	jwt.StandardClaims
}

func GenToken(userID int64, username string) (string, error) {
	c := CustomClaims{
		Username: username,
		UserId:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "bluebell",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(CustomSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
