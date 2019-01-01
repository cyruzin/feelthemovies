package handlers

import (
	"log"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/conn"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db, err = conn.Connect()

// NewRouter initiates the server with the given routes.
// CORS are enabled.
func NewRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/recommendations", getRecommendations).Methods("GET")
	r.HandleFunc("/api/v1/recommendation/{id}", getRecommendation).Methods("GET")
	r.HandleFunc("/api/v1/recommendation", createRecommendation).Methods("POST")
	r.HandleFunc("/api/v1/recommendation/{id}", updateRecommendation).Methods("PUT")
	r.HandleFunc("/api/v1/recommendation/{id}", deleteRecommendation).Methods("DELETE")

	http.Handle("/", r)

	log.Println("MySQL: Connection OK.")
	log.Println("Server: Listening on port 8000.")
	log.Println("You're Good to Go! :)")

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8000", handler))

	return r, nil
}
