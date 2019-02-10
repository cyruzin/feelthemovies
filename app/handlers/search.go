package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/cyruzin/feelthemovies/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"
)

func searchRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := r.URL.Query()
	if len(params) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}
	validate = validator.New()
	err = validate.Var(params["query"][0], "required")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is empty")
		return
	}
	//Redis check
	var rrKey string
	if params["page"] != nil {
		rrKey = fmt.Sprintf(
			"?query=%s?page=%s",
			params["query"][0], params["page"][0],
		)
	} else {
		rrKey = params["query"][0]
	}
	val, err := redisClient.Get(rrKey).Result()
	if err != nil {
		log.Println(err)
	}
	if val != "" {
		w.WriteHeader(200)
		rr := new(model.RecommendationPagination)
		err = helper.UnmarshalBinary([]byte(val), rr)
		json.NewEncoder(w).Encode(rr)
		return
	}
	total, err := db.GetSearchRecommendationTotalRows(params["query"][0]) // total results
	if err != nil {
		log.Println(err)
	}
	var (
		limit       float64 = 10                       // limit per page
		offset      float64                            // offset record
		currentPage float64 = 1                        // current page
		lastPage            = math.Ceil(total / limit) // last page
	)
	// checking if request contains the "page" parameter
	if params["page"] != nil {
		page, err := strconv.ParseFloat(params["page"][0], 64)
		if err != nil {
			log.Println(err)
		}
		if page > currentPage {
			currentPage = page
			offset = (currentPage - 1) * limit
		}
	}
	// End pagination
	search, err := db.SearchRecommendation(offset, limit, params["query"][0])
	if err != nil {
		log.Println(err)
	}
	result := []*model.ResponseRecommendation{}
	for _, r := range search.Data {
		recG, err := db.GetRecommendationGenres(r.ID)
		if err != nil {
			log.Println(err)
		}
		recK, err := db.GetRecommendationKeywords(r.ID)
		if err != nil {
			log.Println(err)
		}
		recFinal := model.ResponseRecommendation{}
		recFinal.Recommendation = r
		recFinal.Genres = recG
		recFinal.Keywords = recK
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

func searchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := r.URL.Query()
	if len(params) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}
	validate = validator.New()
	err = validate.Var(params["query"][0], "required")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is empty")
		return
	}
	search, err := db.SearchUser(params["query"][0])
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}
}

func searchGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := r.URL.Query()
	if len(params) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}
	validate = validator.New()
	err = validate.Var(params["query"][0], "required")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is empty")
		return
	}
	search, err := db.SearchGenre(params["query"][0])
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}
}

func searchKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := r.URL.Query()
	if len(params) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}
	validate = validator.New()
	err = validate.Var(params["query"][0], "required")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is empty")
		return
	}
	search, err := db.SearchKeyword(params["query"][0])
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}
}

func searchSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := r.URL.Query()
	if len(params) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}
	validate = validator.New()
	err = validate.Var(params["query"][0], "required")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is empty")
		return
	}
	search, err := db.SearchSource(params["query"][0])
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}
}
