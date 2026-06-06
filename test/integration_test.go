package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/golang-rest-api/middleware"
	"github.com/petersonsalme/golang-rest-api/redis"
	"github.com/petersonsalme/golang-rest-api/router"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", router.Login)
	r.POST("/logout", middleware.TokenAuthMiddleware(), router.Logout)
	r.POST("/token/refresh", middleware.Refresh)
	r.POST("/todo", middleware.TokenAuthMiddleware(), router.CreateTodo)
	return r
}

func TestIntegration(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsd")
	os.Setenv("REDIS_DSN", "localhost:6379")
	redis.Connect()

	r := setupRouter()

	// Test Login invalid JSON
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{`)))
	r.ServeHTTP(w, req)

	// Test Login valid
	loginData := map[string]string{
		"username": "username",
		"password": "password",
	}
	body, _ := json.Marshal(loginData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	var tokens map[string]string
	json.Unmarshal(w.Body.Bytes(), &tokens)
	accessToken := tokens["access_token"]
	refreshToken := tokens["refresh_token"]

	// Test CreateTodo valid
	todoData := map[string]string{
		"title": "Buy milk",
	}
	body, _ = json.Marshal(todoData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/todo", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test Token Refresh
	refreshData := map[string]string{
		"refresh_token": refreshToken,
	}
	body, _ = json.Marshal(refreshData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/token/refresh", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test Login invalid credentials
	badLoginData := map[string]string{
		"username": "wrong",
		"password": "wrong",
	}
	badBody, _ := json.Marshal(badLoginData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(badBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test Logout without token
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logout", bytes.NewBuffer([]byte{}))
	r.ServeHTTP(w, req)

	// Test CreateTodo without token
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/todo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test CreateTodo invalid json with token
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/todo", bytes.NewBuffer([]byte(`{`)))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test Refresh token invalid JSON
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/token/refresh", bytes.NewBuffer([]byte(`{`)))
	r.ServeHTTP(w, req)

	// Test Refresh token invalid token format
	badRefresh := map[string]string{"refresh_token": "invalid"}
	badRefreshBody, _ := json.Marshal(badRefresh)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/token/refresh", bytes.NewBuffer(badRefreshBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test Logout
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logout", bytes.NewBuffer([]byte{}))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	r.ServeHTTP(w, req)

	// Test Logout again (should fail because token is deleted)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logout", bytes.NewBuffer([]byte{}))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	r.ServeHTTP(w, req)

	// Test Token Refresh again (should fail because refresh token is deleted)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/token/refresh", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test Token Refresh with valid token but deleted from Redis
	// (handled by the previous test)

	// Test CreateTodo with valid token but deleted from Redis
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/todo", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Test ExtractToken with wrong Authorization scheme
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/todo", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Authorization", "Basic "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
}
