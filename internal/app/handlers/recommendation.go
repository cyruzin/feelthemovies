package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	rec, err := model.GetRecommendations(db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rec)
	}
}

func getRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	rec, err := model.GetRecommendation(id, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rec)
	}

}

func createRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var reqRec model.Recommendation

	err = json.NewDecoder(r.Body).Decode(&reqRec)

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
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(rec)
	}
}

func updateRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var reqRec model.Recommendation

	err = json.NewDecoder(r.Body).Decode(&reqRec)

	upRec := model.Recommendation{
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		Status:    reqRec.Status,
		UpdatedAt: time.Now(),
	}

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	rec, err := model.UpdateRecommendation(id, &upRec, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rec)
	}
}

func deleteRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	d, err := model.DeleteRecommendation(id, db)

	if d == 0 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("Something went wrong!")
	}

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Deleted Successfully!")
	}
}
