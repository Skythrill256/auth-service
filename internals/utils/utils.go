package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateJWT(email string, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseJWT(tokenStr string, secret string) (string, error) {
	claims := &jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	email := (*claims)["email"].(string)
	return email, nil
}
