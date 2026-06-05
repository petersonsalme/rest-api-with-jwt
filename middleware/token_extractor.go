package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/petersonsalme/golang-rest-api/model"

	"github.com/golang-jwt/jwt/v5"
)

// ExtractToken should extract token from request
func ExtractToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	regex := regexp.MustCompile("(Bearer\\s)(.*)")
	match := regex.FindStringSubmatch(authorization)

	if len(match) > 0 {
		return match[2]
	}

	return ""
}

// ExtractTokenMetadata ExtractTokenMetadata
func ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &model.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}
