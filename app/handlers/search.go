package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/cyruzin/feelthemovies/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"
)

func searchRecommendation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, "The query field is required", http.StatusBadRequest)
		return
	}
	validate = validator.New()
	if err := validate.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
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
	val, _ := redisClient.Get(rrKey).Result()
	if val != "" {
		rr := &model.RecommendationPagination{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, "Could not unmarshal the payload", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&rr)
		return
	}
	total, err := db.GetSearchRecommendationTotalRows(params["query"][0]) // total results
	if err != nil {
		helper.DecodeError(w, "Could not fetch the search total rows", http.StatusInternalServerError)
		return
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
			helper.DecodeError(w, "Could not parse the page param", http.StatusInternalServerError)
			return
		}
		if page > currentPage {
			currentPage = page
			offset = (currentPage - 1) * limit
		}
	}
	// End pagination
	search, err := db.SearchRecommendation(offset, limit, params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	result := []*model.ResponseRecommendation{}
	for _, r := range search.Data {
		recG, err := db.GetRecommendationGenres(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the genres", http.StatusInternalServerError)
			return
		}
		recK, err := db.GetRecommendationKeywords(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the keywords", http.StatusInternalServerError)
			return
		}
		recFinal := &model.ResponseRecommendation{
			Recommendation: r,
			Genres:         recG,
			Keywords:       recK,
		}
		result = append(result, recFinal)
	}
	resultFinal := &model.RecommendationPagination{
		Data:        result,
		CurrentPage: currentPage,
		LastPage:    lastPage,
		PerPage:     limit,
		Total:       total,
	}
	// Redis set
	rr, err := helper.MarshalBinary(resultFinal)
	if err != nil {
		helper.DecodeError(w, "Could not marshal the payload", http.StatusInternalServerError)
		return
	}
	err = redisClient.Set(rrKey, rr, redisTimeout).Err()
	if err != nil {
		helper.DecodeError(w, "Could not do the set the key", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

func searchUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, "The query field is required", http.StatusBadRequest)
		return
	}
	validate = validator.New()
	if err := validate.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := db.SearchUser(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

func searchGenre(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, "The query field is required", http.StatusBadRequest)
		return
	}
	validate = validator.New()
	if err := validate.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := db.SearchGenre(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

func searchKeyword(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, "The query field is required", http.StatusBadRequest)
		return
	}
	validate = validator.New()
	if err := validate.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := db.SearchKeyword(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

func searchSource(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, "The query field is required", http.StatusBadRequest)
		return
	}
	validate = validator.New()
	if err := validate.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := db.SearchSource(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}
