package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/cyruzin/feelthemovies/internal/app/model"

	"github.com/cyruzin/feelthemovies/internal/app/handlers"
	re "github.com/go-redis/redis"
)

func main() {
	db := database()
	rc := redis()
	mc := model.Connect(db)
	nh := handlers.NewHandler(mc, rc)
	routes(nh)
}

func database() *sql.DB {
	url := fmt.Sprintf(
		"%s:%s@tcp(localhost:3306)/api_feelthemovies?parseTime=true",
		os.Getenv("DBUSER"), os.Getenv("DBPASS"),
	)
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}
	log.Println("MySQL: Connection OK.")
	return db
}

func redis() *re.Client {
	client := re.NewClient(&re.Options{
		Addr:     os.Getenv("REDISADDR"),
		Password: os.Getenv("REDISPASS"),
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Could not open connection to Redis: ", err)
	}
	log.Println("Redis: Connection OK.")
	return client
}

func routes(h *handlers.Setup) {
	r := mux.NewRouter()

	r.Use(h.LoggingMiddleware)
	r.Use(h.AuthMiddleware)

	publicRoutes(r, h)
	authRoutes(r, h)

	http.Handle("/", r)
	handler := cors.AllowAll().Handler(r)

	log.Println("Listening on port: 8000.")
	log.Println("You're good to go! :)")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func publicRoutes(r *mux.Router, h *handlers.Setup) {
	r.HandleFunc("/v1/auth", h.AuthUser).Methods("POST")
}

func authRoutes(r *mux.Router, h *handlers.Setup) {
	r.HandleFunc("/v1/users", h.GetUsers).Methods("GET")
	r.HandleFunc("/v1/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/v1/user/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/v1/user/{id}", h.DeleteUser).Methods("DELETE")

	r.HandleFunc("/v1/recommendations", h.GetRecommendations).Methods("GET")
	r.HandleFunc("/v1/recommendation/{id}", h.GetRecommendation).Methods("GET")
	r.HandleFunc("/v1/recommendation", h.CreateRecommendation).Methods("POST")
	r.HandleFunc("/v1/recommendation/{id}", h.UpdateRecommendation).Methods("PUT")
	r.HandleFunc("/v1/recommendation/{id}", h.DeleteRecommendation).Methods("DELETE")

	r.HandleFunc("/v1/recommendation_items/{id}", h.GetRecommendationItems).Methods("GET")
	r.HandleFunc("/v1/recommendation_item/{id}", h.GetRecommendationItem).Methods("GET")
	r.HandleFunc("/v1/recommendation_item", h.CreateRecommendationItem).Methods("POST")
	r.HandleFunc("/v1/recommendation_item/{id}", h.UpdateRecommendationItem).Methods("PUT")
	r.HandleFunc("/v1/recommendation_item/{id}", h.DeleteRecommendationItem).Methods("DELETE")

	r.HandleFunc("/v1/genres", h.GetGenres).Methods("GET")
	r.HandleFunc("/v1/genre/{id}", h.GetGenre).Methods("GET")
	r.HandleFunc("/v1/genre", h.CreateGenre).Methods("POST")
	r.HandleFunc("/v1/genre/{id}", h.UpdateGenre).Methods("PUT")
	r.HandleFunc("/v1/genre/{id}", h.DeleteGenre).Methods("DELETE")

	r.HandleFunc("/v1/keywords", h.GetKeywords).Methods("GET")
	r.HandleFunc("/v1/keyword/{id}", h.GetKeyword).Methods("GET")
	r.HandleFunc("/v1/keyword", h.CreateKeyword).Methods("POST")
	r.HandleFunc("/v1/keyword/{id}", h.UpdateKeyword).Methods("PUT")
	r.HandleFunc("/v1/keyword/{id}", h.DeleteKeyword).Methods("DELETE")

	r.HandleFunc("/v1/sources", h.GetSources).Methods("GET")
	r.HandleFunc("/v1/source/{id}", h.GetSource).Methods("GET")
	r.HandleFunc("/v1/source", h.CreateSource).Methods("POST")
	r.HandleFunc("/v1/source/{id}", h.UpdateSource).Methods("PUT")
	r.HandleFunc("/v1/source/{id}", h.DeleteSource).Methods("DELETE")

	r.HandleFunc("/v1/search_recommendation", h.SearchRecommendation).Methods("GET")
	r.HandleFunc("/v1/search_user", h.SearchUser).Methods("GET")
	r.HandleFunc("/v1/search_genre", h.SearchGenre).Methods("GET")
	r.HandleFunc("/v1/search_keyword", h.SearchKeyword).Methods("GET")
	r.HandleFunc("/v1/search_source", h.SearchSource).Methods("GET")
}
