package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
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

	m := make(map[int][]int)
	g := []int{11, 4, 13}

	m[3] = g

	//res, err := helper.Sync(m, "genre_recommendation", "recommendation_id", db)
	//res, err := helper.Attach(m, "genre_recommendation", db)
	res, err := helper.Detach(m, "genre_recommendation", "recommendation_id", db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
