package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cyruzin/feelthemovies/internal/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.23.0.2:3306)/api_feelthemovies?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	log.Println("Connection OK!")

	res, err := model.GetRecommendation(1, db)

	// rec := model.Recommendation{
	// 	UserID:    1,
	// 	Title:     "Testando o back em Golang",
	// 	Type:      0,
	// 	Body:      "Corpo da recomendação!",
	// 	Backdrop:  "ahuWashhqQlAsAzz",
	// 	Poster:    "LpOiizzsQRtAmA",
	// 	Status:    0,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// res, err := model.CreateRecommendation(&rec, db)

	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(res, "", "\t")

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", data)

}
