package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/petersonsalme/golang-rest-api/model"
)

// CreateToken creates the token
func CreateToken(userid uint64) (*model.Token, error) {
	token := model.NewToken()

	if err := accessToken(userid, &token); err != nil {
		return nil, err
	}

	if err := refreshToken(userid, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

func accessToken(userid uint64, token *model.Token) error {
	accessSecretValue := os.Getenv("ACCESS_SECRET")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = token.AccessUUID
	atClaims["user_id"] = userid
	atClaims["expires"] = token.AtExpires

	accessTokenProvider := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := accessTokenProvider.SignedString([]byte(accessSecretValue))

	token.AccessToken = accessToken

	return err
}

func refreshToken(userid uint64, token *model.Token) error {
	refreshSecretValue := os.Getenv("REFRESH_SECRET")

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = token.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = token.RtExpires

	refreshTokenProvider := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := refreshTokenProvider.SignedString([]byte(refreshSecretValue))

	token.RefreshToken = refreshToken

	return err
}
