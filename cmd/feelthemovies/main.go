package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/InVisionApp/go-health/handlers"

	"go.uber.org/zap"

	health "github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/checkers"
	redisCheck "github.com/InVisionApp/go-health/checkers/redis"
	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/app/router"

	"github.com/cyruzin/feelthemovies/internal/app/handler"
	re "github.com/go-redis/redis"
	validator "gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate

func main() {
	l, err := zap.NewProduction() // Uber Zap Logger instance.
	if err != nil {
		log.Fatal("Could not initiate the logger")
	}

	db := database()                              // Database instance.
	rc := redis()                                 // Redis client instance.
	mc := model.Connect(db)                       // Passing database instance to the model pkg.
	v = validator.New()                           // Validator instance.
	h := handler.NewHandler(mc, rc, v, l.Sugar()) // Passing instances to the handlers pkg.

	defer l.Sync()
	defer db.Close()
	defer rc.Close()

	healthCheck, err := healthChecks(db) // Health instance.
	if err != nil {
		log.Println("Failed to perform health checks")
	}

	r := router.NewRouter(h, handlers.NewJSONHandlerFunc(healthCheck, nil)) // Passing handlers to the router.

	srv := &http.Server{
		Addr:              ":8000",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           r,
	}

	// Graceful shutdown setup.
	ctx := context.Background()
	gracefulStop := make(chan os.Signal)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		<-gracefulStop
		log.Println("Shutting down the server...")
		srv.Shutdown(ctx) // Shutting down the server gracefully.
		db.Close()        // Closing database and prevent new queries.
		rc.Close()        // Closing redis client.
		l.Sync()          // Flushing buffered log entries.
		os.Exit(0)        // Terminating the app.
	}()

	// Initiating the server.
	log.Println("Listening on port: 8000.")
	log.Println("You're good to go! :)")
	log.Println(srv.ListenAndServe())
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
		Addr:         os.Getenv("REDISADDR"),
		Password:     os.Getenv("REDISPASS"),
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     20,
		PoolTimeout:  30 * time.Second,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Could not open connection to Redis: ", err)
	}
	log.Println("Redis: Connection OK.")
	return client
}

// healthChecks checks the services health periodically.
func healthChecks(db *sql.DB) (*health.Health, error) {
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
			Addr:     os.Getenv("REDISADDR"),
			Password: os.Getenv("REDISPASS"),
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
