package handlers

import (
	"log"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/v1/auth" {
			next.ServeHTTP(w, r)
		} else {
			token := r.Header.Get("Api-Token")
			auth, err := model.CheckAPIToken(token, db)

			if err != nil {
				log.Println(err)
			}

			if auth {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized.", http.StatusForbidden)
			}
		}

	})
}
