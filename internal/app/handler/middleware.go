package handler

import (
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// AuthMiddleware checks if the request contain the Api Token
// on the headers and if it is valid.
func (s *Setup) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")

		token := r.Header.Get("Api-Token")
		if token == "" {
			helper.DecodeError(w, errEmptyToken, http.StatusBadRequest)
			return
		}
		auth, err := s.h.CheckAPIToken(token)
		if err != nil {
			helper.DecodeError(w, errInvalidToken, http.StatusUnauthorized)
			return
		}
		if auth {
			next.ServeHTTP(w, r)
		} else {
			helper.DecodeError(w, errUnauthorized, http.StatusUnauthorized)
			return
		}

	})
}

// NewRelicMiddleware sends data to New Relic servers.
func (s *Setup) NewRelicMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		txn := s.nr.StartTransaction(r.URL.Path, w, r)
		next.ServeHTTP(w, r)
		txn.End()
	}
	return http.HandlerFunc(fn)
}
