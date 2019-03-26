package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"github.com/cyruzin/feelthemovies/internal/app/model"

	"github.com/cyruzin/feelthemovies/internal/app/handler"
	re "github.com/go-redis/redis"
	validator "gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate

func main() {
	db := database() // Database instance.
	defer db.Close()

	rc := redis() // Redis client instance.
	defer rc.Close()

	mc := model.Connect(db) // Passing database instance to the model pkg.

	v = validator.New() // Validator instance.

	h := handler.NewHandler(mc, rc, v) // Passing instances to the handlers pkg.
	router(h)                          // Passing handlers to the router.
}

// Database connection.
func database() *sql.DB {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/api_feelthemovies?parseTime=true",
		os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"),
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

// Redis connection.
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

// All routes setup with CORS and middlewares.
func router(h *handler.Setup) {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Api-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(h.AuthMiddleware)
	r.Use(middleware.Logger)

	publicRoutes(r, h)
	authRoutes(r, h)

	log.Println("Listening on port: 8000.")
	log.Println("You're good to go! :)")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Public routes.
func publicRoutes(r *chi.Mux, h *handler.Setup) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Feel the Movies API V1"))
	})

	r.Post("/auth", h.AuthUser)
}

// Auth routes.
func authRoutes(r *chi.Mux, h *handler.Setup) {

	r.Get("/v1/users", h.GetUsers)
	r.Get("/v1/user/{id}", h.GetUser)
	r.Post("/v1/user", h.CreateUser)
	r.Put("/v1/user/{id}", h.UpdateUser)
	r.Delete("/v1/user/{id}", h.DeleteUser)

	r.Get("/v1/recommendations", h.GetRecommendations)
	r.Get("/v1/recommendation/{id}", h.GetRecommendation)
	r.Post("/v1/recommendation", h.CreateRecommendation)
	r.Put("/v1/recommendation/{id}", h.UpdateRecommendation)
	r.Delete("/v1/recommendation/{id}", h.DeleteRecommendation)

	r.Get("/v1/recommendation_items/{id}", h.GetRecommendationItems)
	r.Get("/v1/recommendation_item/{id}", h.GetRecommendationItem)
	r.Post("/v1/recommendation_item", h.CreateRecommendationItem)
	r.Put("/v1/recommendation_item/{id}", h.UpdateRecommendationItem)
	r.Delete("/v1/recommendation_item/{id}", h.DeleteRecommendationItem)

	r.Get("/v1/genres", h.GetGenres)
	r.Get("/v1/genre/{id}", h.GetGenre)
	r.Post("/v1/genre", h.CreateGenre)
	r.Put("/v1/genre/{id}", h.UpdateGenre)
	r.Delete("/v1/genre/{id}", h.DeleteGenre)

	r.Get("/v1/keywords", h.GetKeywords)
	r.Get("/v1/keyword/{id}", h.GetKeyword)
	r.Post("/v1/keyword", h.CreateKeyword)
	r.Put("/v1/keyword/{id}", h.UpdateKeyword)
	r.Delete("/v1/keyword/{id}", h.DeleteKeyword)

	r.Get("/v1/sources", h.GetSources)
	r.Get("/v1/source/{id}", h.GetSource)
	r.Post("/v1/source", h.CreateSource)
	r.Put("/v1/source/{id}", h.UpdateSource)
	r.Delete("/v1/source/{id}", h.DeleteSource)

	r.HandleFunc("/v1/search_recommendation", h.SearchRecommendation)
	r.HandleFunc("/v1/search_user", h.SearchUser)
	r.HandleFunc("/v1/search_genre", h.SearchGenre)
	r.HandleFunc("/v1/search_keyword", h.SearchKeyword)
	r.HandleFunc("/v1/search_source", h.SearchSource)
}
