package model

import (
	"database/sql"
	"log"

	// MySQL connection driver
	_ "github.com/go-sql-driver/mysql"
)

// Conn type is a struct for connections.
type Conn struct {
	*sql.DB
}

// Connect creates a connection with MySQL database.
func Connect() (*Conn, error) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/api_feelthemovies?parseTime=true")

	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	log.Println("MySQL: Connection OK.")

	return &Conn{db}, err
}
