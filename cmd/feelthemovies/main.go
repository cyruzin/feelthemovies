package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.23.0.2:3306)/api_feelthemovies?parseTime=true")

	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	log.Println("Connection OK!")
}
