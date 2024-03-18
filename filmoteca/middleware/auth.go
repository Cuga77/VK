// middleware/auth.go
package main

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"github.com/Cuga77/filmoteca/models"
)

var jwtKey = []byte("123")

func IsUser(r *http.Request) bool {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		return false
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	return claims.Subject == "user"
}

func IsGuest(r *http.Request) bool {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		return true
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return true
	}

	return claims.Subject == "guest"
}

func CanViewMovie(r *http.Request) bool {
	return IsUser(r) || IsGuest(r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !CanViewMovie(r) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/createActor", AuthMiddleware(http.HandlerFunc(handlers.CreateActorHandler)))
	http.Handle("/getActor", AuthMiddleware(http.HandlerFunc(handlers.GetActorHandler)))
	http.Handle("/updateActor", AuthMiddleware(http.HandlerFunc(handlers.UpdateActorHandler)))
	http.ListenAndServe(":8080", nil)
}
