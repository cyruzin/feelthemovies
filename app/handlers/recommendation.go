package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/gorilla/mux"
)

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := r.URL.Query()

	//Redis check
	var rrKey string
	if params["page"] != nil {
		rrKey = fmt.Sprintf("recommendation?page=%s", params["page"][0])
	} else {
		rrKey = "recommendation"
	}

	val, _ := redisClient.Get(rrKey).Result()

	if val != "" {
		rr := &model.RecommendationPagination{}
		err = helper.UnmarshalBinary([]byte(val), rr)
		json.NewEncoder(w).Encode(rr)
		helper.APIResponse(
			w,
			HTTPOK,
			err,
			"Redis: Could not unmarshal the response",
		)
	}

	// Start pagination
	total, err := db.GetRecommendationTotalRows() // total results
	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Database: Could not count recommendations total rows",
	)

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
			helper.APIResponse(
				w,
				HTTPUnprocessableEntity,
				err,
				"Pagination: Could not parse page to float",
			)
			if page > currentPage {
				currentPage = page
				offset = (currentPage - 1) * limit
			}
		}
	}
	// End pagination

	rec, err := db.GetRecommendations(offset, limit)
	helper.APIResponse(
		w, HTTPUnprocessableEntity,
		err,
		"Database: Could not fetch recommendations",
	)

	result := []*model.ResponseRecommendation{}
	for _, r := range rec.Data {
		recG, err := db.GetRecommendationGenres(r.ID)
		helper.APIResponse(
			w,
			HTTPUnprocessableEntity,
			err,
			"Database: Could not fetch recommendation genres",
		)
		recK, err := db.GetRecommendationKeywords(r.ID)
		helper.APIResponse(
			w,
			HTTPUnprocessableEntity,
			err,
			"Database: Could not fetch recommendation keywords",
		)
		recFinal := model.ResponseRecommendation{
			Recommendation: r,
			Genres:         recG,
			Keywords:       recK,
		}
		result = append(result, &recFinal)
	}

	resultFinal := model.RecommendationPagination{
		Data:        result,
		CurrentPage: currentPage,
		LastPage:    lastPage,
		PerPage:     limit,
		Total:       total,
	}

	// Redis set
	rr, err := helper.MarshalBinary(resultFinal)
	helper.APIResponse(
		w,
		HTTPInternalServerError,
		err,
		"Redis: Could not marshall the response",
	)
	err = redisClient.Set(rrKey, rr, redisTimeout).Err()
	helper.APIResponse(
		w,
		HTTPInternalServerError,
		err,
		"Redis: Could not set the key",
	)

	// Final response
	helper.APIResponse(w, HTTPOK, err, resultFinal)
}

func getRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Parse: Could not parse the ID",
	)

	//Redis check
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := redisClient.Get(rrKey).Result()

	if val != "" {
		rr := new(model.ResponseRecommendation)
		err = helper.UnmarshalBinary([]byte(val), rr)
		helper.APIResponse(
			w,
			HTTPInternalServerError,
			err,
			"Redis: Could not unmarshal the response",
		)
	}

	rec, err := db.GetRecommendation(id)
	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Database: Could fetch the recommendation",
	)

	recG, err := db.GetRecommendationGenres(id)
	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Database: Could not fetch recommendation genres",
	)

	recK, err := db.GetRecommendationKeywords(id)

	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Database: Could not fetch recommendation keywords",
	)

	response := model.ResponseRecommendation{
		Recommendation: rec,
		Genres:         recG,
		Keywords:       recK,
	}

	// Redis set
	rr, err := helper.MarshalBinary(response)
	helper.APIResponse(
		w,
		HTTPInternalServerError,
		err,
		"Redis: Could not marshall the response",
	)

	err = redisClient.Set(rrKey, rr, redisTimeout).Err()
	helper.APIResponse(
		w,
		HTTPInternalServerError,
		err,
		"Redis: Could not set the key",
	)

	helper.APIResponse(w, HTTPOK, err, response)
}

func createRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqRec struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
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
	newRec := model.Recommendation{
		UserID:    int64(reqRec.UserID),
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rec, err := db.CreateRecommendation(&newRec)
	if err != nil {
		log.Println(err)
	}
	// Attaching keyword / genres IDs in their respective pivot tables.
	keywords := make(map[int64][]int)
	genres := make(map[int64][]int)
	keywords[rec.ID] = reqRec.Keywords
	genres[rec.ID] = reqRec.Genres
	_, err = helper.Attach(keywords, "keyword_recommendation", db.DB)
	if err != nil {
		log.Println(err)
	}
	_, err = helper.Attach(genres, "genre_recommendation", db.DB)
	// Redis check
	val, err := redisClient.Get("recommendation").Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		_, err = redisClient.Unlink("recommendation").Result()
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

func updateRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqRec struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
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
	if err != nil {
		log.Println(err)
	}
	// Empty recommendation check
	itemCount, err := db.GetRecommendationItemsTotalRows(id)
	if err != nil {
		log.Println(err)
	}
	if itemCount == 0 && reqRec.Status == 1 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("You can not activate an empty recommendation.")
		return
	}
	rec, err := db.UpdateRecommendation(id, &upRec)
	if err != nil {
		log.Println(err)
	}
	// Syncing keyword / genres IDs in their respective pivot tables.
	keywords := make(map[int64][]int)
	genres := make(map[int64][]int)
	keywords[rec.ID] = reqRec.Keywords
	genres[rec.ID] = reqRec.Genres
	_, err = helper.Sync(keywords, "keyword_recommendation", "recommendation_id", db.DB)
	if err != nil {
		log.Println(err)
	}
	_, err = helper.Sync(genres, "genre_recommendation", "recommendation_id", db.DB)
	// Redis check
	val, err := redisClient.Get("recommendation").Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		_, err = redisClient.Unlink("recommendation").Result()
		if err != nil {
			log.Println(err)
		}
	}
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, err = redisClient.Get(rrKey).Result()
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

func deleteRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Parse: Could not parse the ID",
	)

	// Redis check
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := redisClient.Get(rrKey).Result()

	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		helper.APIResponse(
			w,
			HTTPInternalServerError,
			err,
			"Redis: Could not unlink key",
		)
	}

	val, _ = redisClient.Get("recommendation").Result()

	if val != "" {
		_, err = redisClient.Unlink("recommendation").Result()
		helper.APIResponse(
			w,
			HTTPInternalServerError,
			err,
			"Redis: Could not unlink the key",
		)
	}

	d, err := db.DeleteRecommendation(id)
	helper.APIResponse(
		w,
		HTTPUnprocessableEntity,
		err,
		"Database: Could not delete the recommendation",
	)

	helper.APIResponse(
		w,
		HTTPOK,
		err,
		d,
	)
}
