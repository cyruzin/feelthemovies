package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

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
