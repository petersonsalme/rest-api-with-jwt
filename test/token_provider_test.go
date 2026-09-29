package test

import (
	"testing"

	"github.com/petersonsalme/rest-api-with-jwt/middleware"
	"github.com/petersonsalme/rest-api-with-jwt/model"
)

// TestGivenValidUserIdShouldCreateToken
func TestGivenValidUserIdShouldCreateToken(t *testing.T) {
	token, err := middleware.CreateToken(123)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	validateFields(token, t)
}

func validateFields(token *model.Token, t *testing.T) {
	if token.AccessToken == "" {
		t.Log("AccessToken should not be empty")
		t.Fail()
	}

	if token.RefreshToken == "" {
		t.Log("RefreshToken should not be empty")
		t.Fail()
	}

	if token.AccessUUID == "" {
		t.Log("AccessUUID should not be empty")
		t.Fail()
	}

	if token.RefreshUUID == "" {
		t.Log("RefreshUUID should not be empty")
		t.Fail()
	}

	if token.AtExpires == 0 {
		t.Log("AtExpires should not be 0.")
		t.Fail()
	}

	if token.RtExpires == 0 {
		t.Log("RtExpires should not be 0.")
		t.Fail()
	}
}
