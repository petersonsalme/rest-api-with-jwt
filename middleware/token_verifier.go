package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// VerifyToken verify token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	validationFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	}

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, validationFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid TokenValid
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
