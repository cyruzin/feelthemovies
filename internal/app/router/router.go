package router

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/cors"

	"github.com/cyruzin/feelthemovies/internal/app/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	newrelic "github.com/newrelic/go-agent"
)

// NewRouter has all routes setup with CORS and middlewares.
func NewRouter(h *handler.Setup) *chi.Mux {
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
	r.Use(middleware.Timeout(60 * time.Second))

	authRoutes(r, h)
	publicRoutes(r, h)

	return r
}

// Public routes.
func publicRoutes(r *chi.Mux, h *handler.Setup) {
	app, err := newRelicApp()
	if err != nil {
		log.Println(err)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Feel the Movies API V1"))
	}) // Initial page.

	r.Post("/auth", h.AuthUser) // Authentication end-point.

	r.Get(newrelic.WrapHandleFunc(app, "/v1/recommendations", h.GetRecommendations))
	r.Get("/v1/recommendation/{id}", h.GetRecommendation)

	r.Get(newrelic.WrapHandleFunc(app, "/v1/recommendation_items/{id}", h.GetRecommendationItems))
	r.Get("/v1/recommendation_item/{id}", h.GetRecommendationItem)

	r.Get("/v1/genres", h.GetGenres)
	r.Get("/v1/genre/{id}", h.GetGenre)

	r.Get("/v1/keywords", h.GetKeywords)
	r.Get("/v1/keyword/{id}", h.GetKeyword)

	r.Get("/v1/sources", h.GetSources)
	r.Get("/v1/source/{id}", h.GetSource)

	r.Get(newrelic.WrapHandleFunc(app, "/v1/search_recommendation", h.SearchRecommendation))
	r.Get("/v1/search_genre", h.SearchGenre)
	r.Get("/v1/search_keyword", h.SearchKeyword)
	r.Get("/v1/search_source", h.SearchSource)
}

// Auth routes.
func authRoutes(r *chi.Mux, h *handler.Setup) {
	r.Group(func(r chi.Router) {
		r.Use(h.AuthMiddleware) // Authentication Middleware.

		r.Get("/v1/users", h.GetUsers)
		r.Get("/v1/user/{id}", h.GetUser)
		r.Post("/v1/user", h.CreateUser)
		r.Put("/v1/user/{id}", h.UpdateUser)
		r.Delete("/v1/user/{id}", h.DeleteUser)

		r.Get("/v1/recommendations_admin", h.GetRecommendationsAdmin) // Workaround to list without filter.
		r.Post("/v1/recommendation", h.CreateRecommendation)
		r.Put("/v1/recommendation/{id}", h.UpdateRecommendation)
		r.Delete("/v1/recommendation/{id}", h.DeleteRecommendation)

		r.Post("/v1/recommendation_item", h.CreateRecommendationItem)
		r.Put("/v1/recommendation_item/{id}", h.UpdateRecommendationItem)
		r.Delete("/v1/recommendation_item/{id}", h.DeleteRecommendationItem)

		r.Post("/v1/genre", h.CreateGenre)
		r.Put("/v1/genre/{id}", h.UpdateGenre)
		r.Delete("/v1/genre/{id}", h.DeleteGenre)

		r.Post("/v1/keyword", h.CreateKeyword)
		r.Put("/v1/keyword/{id}", h.UpdateKeyword)
		r.Delete("/v1/keyword/{id}", h.DeleteKeyword)

		r.Post("/v1/source", h.CreateSource)
		r.Put("/v1/source/{id}", h.UpdateSource)
		r.Delete("/v1/source/{id}", h.DeleteSource)

		r.Get("/v1/search_user", h.SearchUser)
	})
}

// New Relic Application instance.
func newRelicApp() (newrelic.Application, error) {
	config := newrelic.NewConfig("Feel the Movies", os.Getenv("NEWRELICKEY"))
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, errors.New("Could not create New Relic Application")
	}
	if err = app.WaitForConnection(time.Duration(10 * time.Second)); err != nil {
		return nil, errors.New("Could not connect to New Relic server")
	}
	log.Println("New Relic: Connection OK.")
	return app, nil
}
