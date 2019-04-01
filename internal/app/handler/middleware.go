package handler

import (
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// JSONMiddleware set content type for all requests.
func (s *Setup) JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware checks if the request contains Bearer Token
// on the headers and if it is valid.
func (s *Setup) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Capturing Authorizathion header.
		tokenHeader := r.Header.Get("Authorization")

		// Checking if the value is empty.
		if tokenHeader == "" {
			helper.DecodeError(w, r, s.l, errEmptyToken, http.StatusBadRequest)
			return
		}

		// Checking if the header contains Bearer string.
		if !strings.Contains(tokenHeader, "Bearer") {
			helper.DecodeError(w, r, s.l, errMalformedToken, http.StatusUnauthorized)
			return
		}

		// Capturing the token.
		jwtString := strings.Split(tokenHeader, "Bearer ")[1]

		// Parsing the token to verify its authenticity.
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTSecret")), nil
		})

		// Returning parsing errors.
		if err != nil {
			helper.DecodeError(w, r, s.l, err.Error(), http.StatusUnauthorized)
			return
		}

		// If the toke is valid.
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			helper.DecodeError(w, r, s.l, errInvalidJWTToken, http.StatusUnauthorized)
			return
		}
	})
}
