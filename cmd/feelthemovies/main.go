package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err = sql.Open("mysql", os.Getenv("MYSQL"))

	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/recommendations", getRecommendations).Methods("GET")
	r.HandleFunc("/api/v1/recommendation/{id}", getRecommendation).Methods("GET")
	r.HandleFunc("/api/v1/recommendation", createRecommendation).Methods("POST")
	r.HandleFunc("/api/v1/recommendation/{id}", updateRecommendation).Methods("PUT")
	r.HandleFunc("/api/v1/recommendation/{id}", deleteRecommendation).Methods("DELETE")

	http.Handle("/", r)

	log.Println("MySQL: Connection OK.")
	log.Println("Server: Listening on port 8000.")
	log.Println("You're Good to Go! :)")

	handler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8000", handler)

	if err != nil {
		log.Fatal(err)
	}

}

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)

	rec, err := model.GetRecommendations(db)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(rec)
}

func getRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	rec, err := model.GetRecommendation(id, db)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(rec)
}

func createRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(201)

	var reqRec model.Recommendation

	err = json.NewDecoder(r.Body).Decode(&reqRec)

	if err != nil {
		log.Fatal(err)
	}

	newRec := model.Recommendation{
		UserID:    reqRec.ID,
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		Status:    reqRec.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rec, err := model.CreateRecommendation(&newRec, db)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(rec)
}

func updateRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)

	var reqRec model.Recommendation

	err = json.NewDecoder(r.Body).Decode(&reqRec)

	if err != nil {
		log.Fatal(err)
	}

	upRec := model.Recommendation{
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		Status:    reqRec.Status,
		UpdatedAt: time.Now(),
	}

	rec, err := model.UpdateRecommendation(reqRec.ID, &upRec, db)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(rec)
}

func deleteRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	_, err = model.DeleteRecommendation(id, db)

	if err != nil {
		log.Fatal(err)
	}

	msg, _ := json.Marshal("Deleted Successfully!")

	w.Write(msg)
}
