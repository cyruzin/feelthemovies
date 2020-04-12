package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	jwt "github.com/dgrijalva/jwt-go"
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
			errhandler.DecodeError(w, r, s.logger, errEmptyToken, http.StatusBadRequest)
			return
		}

		// Checking if the header contains Bearer string and if the token exists.
		if !strings.Contains(tokenHeader, "Bearer") || len(strings.Split(tokenHeader, "Bearer ")) == 1 {
			errhandler.DecodeError(w, r, s.logger, errMalformedToken, http.StatusUnauthorized)
			return
		}

		// Capturing the token.
		jwtString := strings.Split(tokenHeader, "Bearer ")[1]

		// Parsing the token to verify its authenticity.
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWTSECRET")), nil
		})

		// Returning parsing errors.
		if err != nil {
			errhandler.DecodeError(w, r, s.logger, err.Error(), http.StatusUnauthorized)
			return
		}

		// If the token is valid.
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			errhandler.DecodeError(w, r, s.logger, errInvalidJWTToken, http.StatusUnauthorized)
			return
		}
	})
}
