package conn

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	// MySQL connection driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect creates a connection with MySQL database.
func Connect() (*sql.DB, error) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL"))

	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}

	//defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	log.Println("MySQL: Connection OK.")

	return db, err
}
