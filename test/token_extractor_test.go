package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/petersonsalme/rest-api-with-jwt/middleware"
)

// TestGivenValidAuthorizationMustReturnToken
func TestGivenValidAuthorizationMustReturnToken(t *testing.T) {
	expectedToken := "123456789ABC"
	bearerToken := fmt.Sprintf("Bearer %s", expectedToken)

	req := new(http.Request)
	req.Header = map[string][]string{}
	req.Header.Set("Authorization", bearerToken)

	token := middleware.ExtractToken(req)

	if token != expectedToken {
		t.Logf("Token should be equal to [%s].", expectedToken)
		t.Fail()
	}
}

// TestGivenInvalidAuthorizationMustReturnEmpty
func TestGivenInvalidAuthorizationMustReturnEmpty(t *testing.T) {
	req := new(http.Request)
	req.Header = map[string][]string{}

	token := middleware.ExtractToken(req)

	if token != "" {
		t.Log("Token should be empty.")
		t.Fail()
	}
}
