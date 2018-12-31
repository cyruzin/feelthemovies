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

	// r := model.Recommendation{
	// 	UserID:    1,
	// 	Title:     "The Best Soccer Movies of All Time",
	// 	Type:      0,
	// 	Body:      "Only the best movies, no bad movies are allowed!",
	// 	Poster:    "PlQeeRjZkl",
	// 	Backdrop:  "QrTtMklOpz",
	// 	Status:    0,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// rec, err := model.CreateRecommendation(&r, db)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// g := make(map[int64][]int)
	// k := make(map[int64][]int)

	// g[rec.ID] = []int{1, 5, 8}
	// k[rec.ID] = []int{12, 2, 4}

	// _, err = helper.Attach(g, "genre_recommendation", db)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = helper.Attach(k, "keyword_recommendation", db)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := helper.ToJSONIndent(rec)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(data)
}
