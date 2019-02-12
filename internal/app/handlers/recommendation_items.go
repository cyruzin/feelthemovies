package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getRecommendationItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse ID param", http.StatusInternalServerError)
		return
	}

	//Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", id)
	val, _ := redisClient.Get(rrKey).Result()
	if val != "" {
		rr := &model.RecommendationItemFinal{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, "Could not unmarshal the payload", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rr)
		return
	}
	// Redis check end

	rec, err := db.GetRecommendationItems(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation items", http.StatusInternalServerError)
		return
	}

	result := []*model.ResponseRecommendationItem{}
	for _, r := range rec.Data {
		recS, err := db.GetRecommendationItemSources(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the recommendation items sources", http.StatusInternalServerError)
			return
		}

		recFinal := &model.ResponseRecommendationItem{
			RecommendationItem: r,
			Sources:            recS,
		}

		result = append(result, recFinal)
	}

	resultFinal := &model.RecommendationItemFinal{Data: result}

	// Redis set start
	rr, err := helper.MarshalBinary(resultFinal)
	if err != nil {
		helper.DecodeError(w, "Could not unmarshal the payload", http.StatusInternalServerError)
		return
	}

	if err := redisClient.Set(rrKey, rr, redisTimeout).Err(); err != nil {
		helper.DecodeError(w, "Could not set the key", http.StatusInternalServerError)
		return
	}
	// Redis set end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)

}

func getRecommendationItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}

	rec, err := db.GetRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation item", http.StatusInternalServerError)
		return
	}

	recS, err := db.GetRecommendationItemSources(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation item source", http.StatusInternalServerError)
		return
	}

	response := &model.ResponseRecommendationItem{
		RecommendationItem: rec,
		Sources:            recS,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createRecommendationItem(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, "Could not decode request body", http.StatusInternalServerError)
		return
	}

	validate = validator.New()
	if err := validate.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	// Parsing string to time.Time
	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)
	if err != nil {
		helper.DecodeError(w, "Could not parse the date", http.StatusInternalServerError)
		return
	}

	newRec := &model.RecommendationItem{
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

	rec, err := db.CreateRecommendationItem(newRec)
	if err != nil {
		helper.DecodeError(w, "Could not create the recommendation item", http.StatusInternalServerError)
		return
	}

	// Attaching sources
	sources := make(map[int64][]int)
	sources[rec.ID] = reqRec.Sources
	err = helper.Attach(sources, "recommendation_item_source", db.DB)
	if err != nil {
		helper.DecodeError(w, "Could not attach the recommendation item sources", http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, _ := redisClient.Get(rrKey).Result()
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&rec)

}

func updateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}

	validate = validator.New()
	if err := validate.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)
	if err != nil {
		helper.DecodeError(w, "Could not parse the date", http.StatusInternalServerError)
		return
	}

	upRec := &model.RecommendationItem{
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
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}

	rec, err := db.UpdateRecommendationItem(id, upRec)
	if err != nil {
		helper.DecodeError(w, "Could not update the recommendation item", http.StatusInternalServerError)
		return
	}

	// Syncing sources
	sources := make(map[int64][]int)
	sources[rec.ID] = reqRec.Sources
	err = helper.Sync(sources, "recommendation_item_source", "recommendation_item_id", db.DB)
	if err != nil {
		helper.DecodeError(w, "Could not sync the recommendation item sources", http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, _ := redisClient.Get(rrKey).Result()
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&rec)
}

func deleteRecommendationItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}

	rec, err := db.GetRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, "Could not get fetch the recommendation item", http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, _ := redisClient.Get(rrKey).Result()
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}

	err = db.DeleteRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, "Could not delete the recommendation item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation item deleted successfully!"})
}
