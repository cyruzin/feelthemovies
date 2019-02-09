package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/cyruzin/feelthemovies/pkg/helper"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getRecommendationItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	//Redis check
	rrKey := fmt.Sprintf("recommendation_items-%d", id)
	val, err := redisClient.Get(rrKey).Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		w.WriteHeader(200)
		rr := new(model.RecommendationItemFinal)
		err = helper.UnmarshalBinary([]byte(val), rr)
		json.NewEncoder(w).Encode(rr)
		return
	}
	rec, err := db.GetRecommendationItems(id)
	if err != nil {
		log.Println(err)
	}
	result := []*model.ResponseRecommendationItem{}
	for _, r := range rec.Data {
		recS, err := db.GetRecommendationItemSources(r.ID)
		if err != nil {
			log.Println(err)
		}
		recFinal := model.ResponseRecommendationItem{}
		recFinal.RecommendationItem = r
		recFinal.Sources = recS
		result = append(result, &recFinal)
	}
	resultFinal := model.RecommendationItemFinal{Data: result}
	// Redis set
	rr, err := helper.MarshalBinary(resultFinal)
	if err != nil {
		log.Println(err)
	}
	err = redisClient.Set(rrKey, rr, redisTimeout).Err()
	if err != nil {
		log.Println(err)
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
	if err != nil {
		log.Println(err)
	}
	rec, err := db.GetRecommendationItem(id)
	if err != nil {
		log.Println(err)
	}
	recS, err := db.GetRecommendationItemSources(id)
	if err != nil {
		log.Println(err)
	}
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
	err := json.NewDecoder(r.Body).Decode(&reqRec)
	if err != nil {
		log.Println(err)
	}
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
	if err != nil {
		log.Println(err)
	}
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
	rec, err := db.CreateRecommendationItem(&newRec)
	if err != nil {
		log.Println(err)
	}
	// Attaching sources IDs in its respective pivot table.
	sources := make(map[int64][]int)
	sources[rec.ID] = reqRec.Sources
	_, err = helper.Attach(sources, "recommendation_item_source", db.DB)
	// Redis check
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, err := redisClient.Get(rrKey).Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			log.Println(err)
		}
	}
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
	err := json.NewDecoder(r.Body).Decode(&reqRec)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqRec)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)
	if err != nil {
		log.Println(err)
	}
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
	if err != nil {
		log.Println(err)
	}
	rec, err := db.UpdateRecommendationItem(id, &upRec)
	if err != nil {
		log.Println(err)
	}
	// Syncing sources IDs in its respective pivot table.
	sources := make(map[int64][]int)
	sources[rec.ID] = reqRec.Sources
	_, err = helper.Sync(sources, "recommendation_item_source", "recommendation_item_id", db.DB)
	// Redis check
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, err := redisClient.Get(rrKey).Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			log.Println(err)
		}
	}
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
	if err != nil {
		log.Println(err)
	}
	rec, err := db.GetRecommendationItem(id)
	if err != nil {
		log.Println(err)
	}
	// Redis check
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, err := redisClient.Get(rrKey).Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			log.Println(err)
		}
	}
	d, err := db.DeleteRecommendationItem(id)
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
