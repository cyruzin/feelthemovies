package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/app/model"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Initializing database connection.
var db, err = model.Connect()

// Initializing Redis.
var redisClient = model.Redis()

// Redis expiration time.
const redisTimeout = time.Duration(5 * time.Minute)

// Validator instance
var validate *validator.Validate

// NewRouter initiates the server with the given routes.
// CORS are enabled.
func NewRouter() *mux.Router {
	defer db.Close()
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(authMiddleware)
	publicRoutes(r)
	authRoutes(r)
	http.Handle("/", r)
	handler := cors.AllowAll().Handler(r)
	log.Println("Listening on port: 8000.")
	log.Println("You're good to go! :)")
	log.Fatal(http.ListenAndServe(":8000", handler))
	return r
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
	r.HandleFunc("/v1/search_recommendation", searchRecommendation).Methods("GET")
	r.HandleFunc("/v1/search_user", searchUser).Methods("GET")
	r.HandleFunc("/v1/search_genre", searchGenre).Methods("GET")
	r.HandleFunc("/v1/search_keyword", searchKeyword).Methods("GET")
	r.HandleFunc("/v1/search_source", searchSource).Methods("GET")
}
