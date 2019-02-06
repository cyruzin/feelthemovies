package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
		"%s:%s@tcp(localhost:3306)/api_feelthemovies?parseTime=true",
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
