package handlers

import (
	"log"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/conn"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// TODO: Implement form validation
// TODO: Implement Attach / Sync functions on handlers

// Initializing database connection.
var db, err = conn.Connect()

// NewRouter initiates the server with the given routes.
// CORS are enabled.
func NewRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(authMiddleware)

	publicRoutes(r)
	authRoutes(r)

	http.Handle("/", r)

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8000", handler))

	return r, nil

}

func publicRoutes(r *mux.Router) {
	r.HandleFunc("/v1/auth", authUser).Methods("POST")
}

func authRoutes(r *mux.Router) {
	r.HandleFunc("/v1/users", getUsers).Methods("GET")
	r.HandleFunc("/v1/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/v1/user", createUser).Methods("POST")
	r.HandleFunc("/v1/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/v1/user/{id}", deleteUser).Methods("DELETE")

	r.HandleFunc("/v1/recommendations", getRecommendations).Methods("GET")
	r.HandleFunc("/v1/recommendation/{id}", getRecommendation).Methods("GET")
	r.HandleFunc("/v1/recommendation", createRecommendation).Methods("POST")
	r.HandleFunc("/v1/recommendation/{id}", updateRecommendation).Methods("PUT")
	r.HandleFunc("/v1/recommendation/{id}", deleteRecommendation).Methods("DELETE")

	r.HandleFunc("/v1/recommendation_items/{id}", getRecommendationItems).Methods("GET")
	r.HandleFunc("/v1/recommendation_item/{id}", getRecommendationItem).Methods("GET")
	r.HandleFunc("/v1/recommendation_item", createRecommendationItem).Methods("POST")
	r.HandleFunc("/v1/recommendation_item/{id}", updateRecommendationItem).Methods("PUT")
	r.HandleFunc("/v1/recommendation_item/{id}", deleteRecommendationItem).Methods("DELETE")

	r.HandleFunc("/v1/genres", getGenres).Methods("GET")
	r.HandleFunc("/v1/genre/{id}", getGenre).Methods("GET")
	r.HandleFunc("/v1/genre", createGenre).Methods("POST")
	r.HandleFunc("/v1/genre/{id}", updateGenre).Methods("PUT")
	r.HandleFunc("/v1/genre/{id}", deleteGenre).Methods("DELETE")

	r.HandleFunc("/v1/keywords", getKeywords).Methods("GET")
	r.HandleFunc("/v1/keyword/{id}", getKeyword).Methods("GET")
	r.HandleFunc("/v1/keyword", createKeyword).Methods("POST")
	r.HandleFunc("/v1/keyword/{id}", updateKeyword).Methods("PUT")
	r.HandleFunc("/v1/keyword/{id}", deleteKeyword).Methods("DELETE")

	r.HandleFunc("/v1/sources", getSources).Methods("GET")
	r.HandleFunc("/v1/source/{id}", getSource).Methods("GET")
	r.HandleFunc("/v1/source", createSource).Methods("POST")
	r.HandleFunc("/v1/source/{id}", updateSource).Methods("PUT")
	r.HandleFunc("/v1/source/{id}", deleteSource).Methods("DELETE")
}
