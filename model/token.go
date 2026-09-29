package model

import (
	"time"

	"github.com/twinj/uuid"
)

// Token describes all token's information
type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// NewToken creates new Token instance
func NewToken() Token {
	token := Token{}

	token.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	token.AccessUUID = uuid.NewV4().String()

	token.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.RefreshUUID = uuid.NewV4().String()

	return token
}
