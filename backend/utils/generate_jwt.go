package utils

import (
	"backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.AppConfig.JWT.ExpireHours) * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(), // 失效时间7天后
	})

	jwtKey := []byte(config.AppConfig.JWT.Secret)
	signedToken, err := token.SignedString(jwtKey) // 签名
	return signedToken, err
}
