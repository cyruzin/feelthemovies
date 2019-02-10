package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	// MySQL connection driver
	_ "github.com/go-sql-driver/mysql"
)

// Conn type is a struct for connections.
type Conn struct {
	*sql.DB
}

// Connect creates a connection with MySQL database.
func Connect() (*Conn, error) {
	url := fmt.Sprintf(
		"%s:%s@tcp(http://feelthemovies.com.br:3306)/api_feelthemovies?parseTime=true",
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
	return &Conn{db}, nil
}

// Redis func creates a connection with the Redis server.
func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
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
