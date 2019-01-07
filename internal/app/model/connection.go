package model

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	// MySQL connection driver
	_ "github.com/go-sql-driver/mysql"
)

// Conn type is a struct for connections.
type Conn struct {
	*sql.DB
}

// Connect creates a connection with MySQL database.
func Connect() (*Conn, error) {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/cyruzin/feelthemovies/.env"))

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL"))

	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	log.Println("MySQL: Connection OK.")

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	return &Conn{db}, err
}
