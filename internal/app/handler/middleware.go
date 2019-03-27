package handler

import (
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// JSONMiddleware set content type for all requests.
func (s *Setup) JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware checks if the request contain the Api Token
// on the headers and if it is valid.
func (s *Setup) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Api-Token")
		if token == "" {
			helper.DecodeError(w, nil, errEmptyToken, http.StatusBadRequest)
			return
		}
		auth, err := s.h.CheckAPIToken(token)
		if err != nil {
			helper.DecodeError(w, err, errInvalidToken, http.StatusUnauthorized)
			return
		}
		if auth {
			next.ServeHTTP(w, r)
		} else {
			helper.DecodeError(w, err, errUnauthorized, http.StatusUnauthorized)
			return
		}

	})
}
