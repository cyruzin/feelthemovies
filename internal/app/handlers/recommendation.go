package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	// Start pagination
	params := r.URL.Query()

	total, err := model.GetRecommendationTotalRows(db) // total results
	var (
		limit       float64 = 10                       // limit per page
		offset      float64                            // offset record
		currentPage float64 = 1                        // current page
		lastPage            = math.Ceil(total / limit) // last page
	)

	// checking if request contains the "page" parameter
	if len(params) > 0 {
		if params["page"][0] != "" {
			page, err := strconv.ParseFloat(params["page"][0], 64)

			if err != nil {
				log.Println(err)
			}

			if page > currentPage {
				currentPage = page
				offset = (currentPage - 1) * limit
			}
		}
	}

	// End pagination

	rec, err := model.GetRecommendations(offset, limit, db)

	result := []*model.ResponseRecommendation{}

	for _, r := range rec.Data {
		recG, err := model.GetRecommendationGenres(r.ID, db)
		recK, err := model.GetRecommendationKeywords(r.ID, db)

		if err != nil {
			log.Println(err)
		}

		recFinal := model.ResponseRecommendation{}

		recFinal.Recommendation = r
		recFinal.Genres = recG
		recFinal.Keywords = recK

		result = append(result, &recFinal)
	}

	resultFinal := model.RecommendationPagination{}

	resultFinal.Data = result
	resultFinal.CurrentPage = currentPage
	resultFinal.LastPage = lastPage
	resultFinal.PerPage = limit
	resultFinal.Total = total

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resultFinal)
	}
}

func getRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	rec, err := model.GetRecommendation(id, db)
	recG, err := model.GetRecommendationGenres(id, db)
	recK, err := model.GetRecommendationKeywords(id, db)

	response := model.ResponseRecommendation{}

	response.Recommendation = rec
	response.Genres = recG
	response.Keywords = recK

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}

}

func createRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var reqRec struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
	}

	err = json.NewDecoder(r.Body).Decode(&reqRec)

	validate = validator.New()
	err = validate.Struct(reqRec)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}

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

	var reqRec struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
	}

	err = json.NewDecoder(r.Body).Decode(&reqRec)

	validate = validator.New()
	err = validate.Struct(reqRec)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
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