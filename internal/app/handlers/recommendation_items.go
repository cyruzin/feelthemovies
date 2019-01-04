package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getRecommendationItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	rec, err := model.GetRecommendationItems(id, db)

	result := []*model.ResponseRecommendationItem{}

	for _, r := range rec.Data {
		recS, err := model.GetRecommendationItemSources(r.ID, db)

		if err != nil {
			log.Println(err)
		}

		recFinal := model.ResponseRecommendationItem{}

		recFinal.RecommendationItem = r
		recFinal.Sources = recS

		result = append(result, &recFinal)
	}

	resultFinal := struct {
		Data []*model.ResponseRecommendationItem `json:"data"`
	}{
		result,
	}

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resultFinal)
	}
}

func getRecommendationItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	rec, err := model.GetRecommendationItem(id, db)
	recS, err := model.GetRecommendationItemSources(id, db)

	response := model.ResponseRecommendationItem{}

	response.RecommendationItem = rec
	response.Sources = recS

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}

}

func createRecommendationItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var reqRec struct {
		*model.RecommendationItem
		Sources []int  `json:"sources" validate:"required"`
		Year    string `json:"year" validate:"required"`
	}

	err = json.NewDecoder(r.Body).Decode(&reqRec)
	validate = validator.New()
	err = validate.Struct(reqRec)

	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}

	// Parsing string to time.Time
	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)

	newRec := model.RecommendationItem{
		RecommendationID: reqRec.RecommendationID,
		Name:             reqRec.Name,
		TMDBID:           reqRec.TMDBID,
		Year:             yearParsed,
		Overview:         reqRec.Overview,
		Poster:           reqRec.Poster,
		Backdrop:         reqRec.Backdrop,
		Trailer:          reqRec.Trailer,
		Commentary:       reqRec.Commentary,
		MediaType:        reqRec.MediaType,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	rec, err := model.CreateRecommendationItem(&newRec, db)

	// Attaching sources IDs in its respective pivot table.
	sources := make(map[int64][]int)

	sources[rec.ID] = reqRec.Sources

	_, err = helper.Attach(sources, "recommendation_item_source", db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(rec)
	}
}

func updateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var reqRec struct {
		*model.RecommendationItem
		Sources []int  `json:"sources" validate:"required"`
		Year    string `json:"year" validate:"required"`
	}

	err = json.NewDecoder(r.Body).Decode(&reqRec)

	validate = validator.New()
	err = validate.Struct(reqRec)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}

	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)

	upRec := model.RecommendationItem{
		Name:       reqRec.Name,
		TMDBID:     reqRec.TMDBID,
		Year:       yearParsed,
		Overview:   reqRec.Overview,
		Poster:     reqRec.Poster,
		Backdrop:   reqRec.Backdrop,
		Trailer:    reqRec.Trailer,
		Commentary: reqRec.Commentary,
		MediaType:  reqRec.MediaType,
		UpdatedAt:  time.Now(),
	}

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	rec, err := model.UpdateRecommendationItem(id, &upRec, db)

	// Syncing sources IDs in its respective pivot table.
	sources := make(map[int64][]int)

	sources[rec.ID] = reqRec.Sources

	_, err = helper.Sync(sources, "recommendation_item_source", "recommendation_item_id", db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rec)
	}
}

func deleteRecommendationItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	d, err := model.DeleteRecommendationItem(id, db)

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
