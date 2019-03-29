package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

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
	defer l.Sync()

	db := database() // Database instance.
	defer db.Close()

	rc := redis() // Redis client instance.
	defer rc.Close()

	mc := model.Connect(db) // Passing database instance to the model pkg.

	v = validator.New() // Validator instance.

	h := handler.NewHandler(mc, rc, v, l.Sugar()) // Passing instances to the handlers pkg.
	r := router.NewRouter(h)                      // Passing handlers to the router.

	log.Println("Listening on port: 8000.")
	log.Println("You're good to go! :)")
	log.Fatal(http.ListenAndServe(":8000", r)) // Initiating the server.
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
