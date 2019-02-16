package handlers

import (
	"log"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// LoggingMiddleware ...
func (s *Setup) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware ...
func (s *Setup) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")
		if r.RequestURI == "/v1/auth" {
			next.ServeHTTP(w, r)
		} else {
			token := r.Header.Get("Api-Token")
			auth, err := s.h.CheckAPIToken(token)

			if err != nil {
				helper.DecodeError(w, "Invalid API Token", http.StatusUnauthorized)
				return
			}

			if auth {
				next.ServeHTTP(w, r)
			} else {
				helper.DecodeError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
	})
}
