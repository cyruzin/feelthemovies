package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/InVisionApp/go-health/handlers"
	"github.com/jmoiron/sqlx"

	health "github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/checkers"
	redisCheck "github.com/InVisionApp/go-health/checkers/redis"
	"github.com/cyruzin/feelthemovies/internal/app/config"
	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/app/router"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"

	"github.com/cyruzin/feelthemovies/internal/app/handler"
	re "github.com/go-redis/redis"
	validator "gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate

func main() {
	l, err := logger.Init() // Uber Zap Logger instance.
	if err != nil {
		log.Fatal("Could not initiate the logger: " + err.Error())
	}

	cfg, err := config.Load() // Loading environment variables.
	if err != nil {
		log.Fatal(err.Error())
	}

	db := database(cfg)                   // Database instance.
	rc := redis(cfg)                      // Redis client instance.
	mc := model.Connect(db)               // Passing database instance to the model pkg.
	v = validator.New()                   // Validator instance.
	h := handler.NewHandler(mc, rc, v, l) // Passing instances to the handlers pkg.

	defer l.Sync()
	defer db.Close()
	defer rc.Close()

	healthCheck, err := healthChecks(cfg, db) // Health instance.
	if err != nil {
		log.Println("Failed to perform health checks")
	}

	r := router.NewRouter(h, handlers.NewJSONHandlerFunc(healthCheck, nil)) // Passing handlers to the router.

	srv := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           r,
	}

	// Graceful shutdown setup.
	idleConnsClosed := make(chan struct{})

	go func() {
		gracefulStop := make(chan os.Signal, 1)
		signal.Notify(gracefulStop, os.Interrupt)
		<-gracefulStop

		log.Println("Shutting down the server...")
		if err := srv.Shutdown(context.Background()); err != nil { // Shutting down the server gracefully.
			log.Printf("HTTP server Shutdown: %v", err) // Error from closing listeners, or context timeout.
		}
		close(idleConnsClosed) // Closing channel.
	}()

	// Initiating the server.
	log.Println("Listening on port: " + cfg.ServerPort)
	log.Println("You're good to go! :)")
	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err) // Error starting or closing listener.
	}
	<-idleConnsClosed
}

// Database connection.
func database(cfg *config.Config) *sqlx.DB {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost,
		cfg.DBPort, cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", url)
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
func redis(cfg *config.Config) *re.Client {
	client := re.NewClient(&re.Options{
		Addr:         cfg.RedisAddress,
		Password:     cfg.RedisPass,
		DB:           0,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		PoolSize:     20,
		PoolTimeout:  10 * time.Second,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Could not open connection to Redis: ", err)
	}
	log.Println("Redis: Connection OK.")
	return client
}

// healthChecks checks the services health periodically.
func healthChecks(cfg *config.Config, db *sqlx.DB) (*health.Health, error) {
	h := health.New()
	h.DisableLogging()

	mysqlDB, err := checkers.NewSQL(&checkers.SQLConfig{
		Pinger: db,
	})
	if err != nil {
		return nil, err
	}

	redisDB, err := redisCheck.NewRedis(&redisCheck.RedisConfig{
		Auth: &redisCheck.RedisAuthConfig{
			Addr:     cfg.RedisAddress,
			Password: cfg.RedisPass,
			DB:       0,
		},
		Ping: true,
	})
	if err != nil {
		return nil, err
	}

	if err = h.AddChecks([]*health.Config{
		{
			Name:     "feelthemovies-database",
			Checker:  mysqlDB,
			Interval: time.Duration(3) * time.Second,
			Fatal:    true,
		},
		{
			Name:     "feelthemovies-redis",
			Checker:  redisDB,
			Interval: time.Duration(3) * time.Second,
			Fatal:    true,
		},
	}); err != nil {
		return nil, err
	}

	if err := h.Start(); err != nil {
		return nil, err
	}

	return h, nil
}
