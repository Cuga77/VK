package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestIsUser(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: "user",
	})

	tokenStr, _ := token.SignedString(jwtKey)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)

	if !IsUser(req) {
		t.Errorf("IsUser() = false; want true")
	}
}

func TestIsGuest(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: "guest",
	})

	tokenStr, _ := token.SignedString(jwtKey)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)

	if !IsGuest(req) {
		t.Errorf("IsGuest() = false; want true")
	}
}

func TestCanViewMovie(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: "user",
	})

	tokenStr, _ := token.SignedString(jwtKey)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)

	if !CanViewMovie(req) {
		t.Errorf("CanViewMovie() = false; want true")
	}
}

func TestAuthMiddleware(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: "user",
	})

	tokenStr, _ := token.SignedString(jwtKey)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)

	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}