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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	loggerInstance, err := logger.Init() // Uber Zap Logger instance.
	if err != nil {
		log.Fatal("Could not initiate the logger: " + err.Error())
	}

	cfg, err := config.Load() // Loading environment variables.
	if err != nil {
		log.Fatal(err.Error())
	}

	databaseInstance := database(ctx, cfg)           // Database instance.
	redisInstance := redis(ctx, cfg)                 // Redis client instance.
	modelInstance := model.Connect(databaseInstance) // Passing database instance to the model pkg.
	validatorInstance := validator.New()             // Validator instance.
	handlersInstance := handler.NewHandler(
		modelInstance,
		redisInstance,
		validatorInstance,
		loggerInstance,
	) // Passing instances to the handlers pkg.

	defer loggerInstance.Sync()
	defer databaseInstance.Close()
	defer redisInstance.Close()

	healthCheck, err := healthChecks(cfg, databaseInstance) // Health instance.
	if err != nil {
		log.Println("Failed to perform health checks")
	}

	r := router.NewRouter(
		handlersInstance,
		handlers.NewJSONHandlerFunc(healthCheck, nil),
	) // Passing handlers to the router.

	srv := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      20 * time.Second,
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
func database(ctx context.Context, cfg *config.Config) *sqlx.DB {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost,
		cfg.DBPort, cfg.DBName,
	)

	db, err := sqlx.ConnectContext(ctx, "mysql", url)
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
func redis(ctx context.Context, cfg *config.Config) *re.Client {
	client := re.NewClient(&re.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPass,
		DB:       0,
	})

	_, err := client.WithContext(ctx).Ping().Result()
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
