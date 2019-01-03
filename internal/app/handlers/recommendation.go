package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"

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
		UserID:    reqRec.UserID,
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

	// Attaching keyword / genres IDs in their respective pivot tables.
	keywords := make(map[int64][]int)
	genres := make(map[int64][]int)

	keywords[rec.ID] = reqRec.Keywords
	genres[rec.ID] = reqRec.Genres

	_, err = helper.Attach(keywords, "keyword_recommendation", db)
	_, err = helper.Attach(genres, "genre_recommendation", db)

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

	// Syncing keyword / genres IDs in their respective pivot tables.
	keywords := make(map[int64][]int)
	genres := make(map[int64][]int)

	keywords[rec.ID] = reqRec.Keywords
	genres[rec.ID] = reqRec.Genres

	_, err = helper.Sync(keywords, "keyword_recommendation", "recommendation_id", db)
	_, err = helper.Sync(genres, "genre_recommendation", "recommendation_id", db)

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

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else if d == 0 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Deleted Successfully!")
	}
}
