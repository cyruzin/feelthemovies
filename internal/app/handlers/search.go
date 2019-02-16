package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"
)

// SearchRecommendation ...
func (s *Setup) SearchRecommendation(w http.ResponseWriter, r *http.Request) {
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
	val, _ := s.rc.Get(rrKey).Result()
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
	total, err := s.h.GetSearchRecommendationTotalRows(params["query"][0]) // total results
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
	search, err := s.h.SearchRecommendation(offset, limit, params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	result := []*model.RecommendationResponse{}
	for _, r := range search.Data {
		recG, err := s.h.GetRecommendationGenres(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the genres", http.StatusInternalServerError)
			return
		}
		recK, err := s.h.GetRecommendationKeywords(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the keywords", http.StatusInternalServerError)
			return
		}
		recFinal := &model.RecommendationResponse{
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
	err = s.rc.Set(rrKey, rr, redisTimeout).Err()
	if err != nil {
		helper.DecodeError(w, "Could not do the set the key", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

// SearchUser ...
func (s *Setup) SearchUser(w http.ResponseWriter, r *http.Request) {
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
	search, err := s.h.SearchUser(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchGenre ...
func (s *Setup) SearchGenre(w http.ResponseWriter, r *http.Request) {
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
	search, err := s.h.SearchGenre(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchKeyword ...
func (s *Setup) SearchKeyword(w http.ResponseWriter, r *http.Request) {
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
	search, err := s.h.SearchKeyword(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchSource ...
func (s *Setup) SearchSource(w http.ResponseWriter, r *http.Request) {
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
	search, err := s.h.SearchSource(params["query"][0])
	if err != nil {
		helper.DecodeError(w, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}
