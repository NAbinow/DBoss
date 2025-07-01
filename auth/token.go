package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("this is a secret")

func Create_JWT(email_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email_id": email_id,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenstring, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return tokenstring, nil

}

func Verify_JWT(token_string string) (string, error) {
	token, err := jwt.Parse(token_string, func(t *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}
	return (token.Claims.(jwt.MapClaims))["email_id"].(string), nil
}
