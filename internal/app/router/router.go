package router

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/controllers"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	newrelic "github.com/newrelic/go-agent"
)

// New has all routes setup with CORS and middlewares.
func New(
	c *controllers.Setup,
	healthHandler http.Handler,
	logger *logger.Logger) *chi.Mux {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(
		c.LoggerMiddleware,
		render.SetContentType(render.ContentTypeJSON),
		cors.Handler,
	)

	authRoutes(router, c)
	publicRoutes(router, c, logger)
	router.Handle("/healthcheck", healthHandler)

	return router
}

// Public routes.
func publicRoutes(
	r *chi.Mux,
	c *controllers.Setup,
	logger *logger.Logger,
) {
	app, err := newRelicApp(logger)
	if err != nil {
		logger.Warn(err)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("https://github.com/cyruzin/feelthemovies"))
	})

	r.Post("/auth", c.AuthUser)

	r.Get(newrelic.WrapHandleFunc(app, "/v1/recommendations", c.GetRecommendations))
	r.Get("/v1/recommendation/{id}", c.GetRecommendation)

	r.Get(newrelic.WrapHandleFunc(app, "/v1/recommendation_items/{id}", c.GetRecommendationItems))
	r.Get("/v1/recommendation_item/{id}", c.GetRecommendationItem)

	r.Get("/v1/genres", c.GetGenres)
	r.Get("/v1/genre/{id}", c.GetGenre)

	r.Get("/v1/keywords", c.GetKeywords)
	r.Get("/v1/keyword/{id}", c.GetKeyword)

	r.Get("/v1/sources", c.GetSources)
	r.Get("/v1/source/{id}", c.GetSource)

	r.Get(newrelic.WrapHandleFunc(app, "/v1/search_recommendation", c.SearchRecommendation))
	r.Get("/v1/search_genre", c.SearchGenre)
	r.Get("/v1/search_keyword", c.SearchKeyword)
	r.Get("/v1/search_source", c.SearchSource)
}

// Auth routes.
func authRoutes(
	r *chi.Mux,
	c *controllers.Setup,
) {
	r.Group(func(r chi.Router) {
		r.Use(c.AuthMiddleware)

		r.Get("/v1/users", c.GetUsers)
		r.Get("/v1/user/{id}", c.GetUser)
		r.Post("/v1/user", c.CreateUser)
		r.Put("/v1/user/{id}", c.UpdateUser)
		r.Delete("/v1/user/{id}", c.DeleteUser)

		r.Get("/v1/recommendations_admin", c.GetRecommendationsAdmin)
		r.Get("/v1/recommendation_genres/{id}", c.GetRecommendationGenres)
		r.Get("/v1/recommendation_keywords/{id}", c.GetRecommendationKeywords)
		r.Post("/v1/recommendation", c.CreateRecommendation)
		r.Put("/v1/recommendation/{id}", c.UpdateRecommendation)
		r.Delete("/v1/recommendation/{id}", c.DeleteRecommendation)

		r.Get("/v1/recommendation_item_sources/{id}", c.GetRecommendationItemSources)
		r.Post("/v1/recommendation_item", c.CreateRecommendationItem)
		r.Put("/v1/recommendation_item/{id}", c.UpdateRecommendationItem)
		r.Delete("/v1/recommendation_item/{id}", c.DeleteRecommendationItem)

		r.Post("/v1/genre", c.CreateGenre)
		r.Put("/v1/genre/{id}", c.UpdateGenre)
		r.Delete("/v1/genre/{id}", c.DeleteGenre)

		r.Post("/v1/keyword", c.CreateKeyword)
		r.Put("/v1/keyword/{id}", c.UpdateKeyword)
		r.Delete("/v1/keyword/{id}", c.DeleteKeyword)

		r.Post("/v1/source", c.CreateSource)
		r.Put("/v1/source/{id}", c.UpdateSource)
		r.Delete("/v1/source/{id}", c.DeleteSource)

		r.Get("/v1/search_user", c.SearchUser)
	})
}

// New Relic Application instance.
func newRelicApp(logger *logger.Logger) (newrelic.Application, error) {
	config := newrelic.NewConfig("Feel the Movies", os.Getenv("NEWRELICKEY"))

	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, errors.New("Could not create New Relic Application")
	}

	if err = app.WaitForConnection(time.Duration(10 * time.Second)); err != nil {
		return nil, errors.New("Could not connect to New Relic server")
	}

	logger.Info("New Relic: Connection OK.")

	return app, nil
}
