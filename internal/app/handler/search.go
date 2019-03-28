package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// SearchRecommendation searches for recommendations.
func (s *Setup) SearchRecommendation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.l, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.v.Var(params["query"][0], "required"); err != nil {
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
			helper.DecodeError(w, r, s.l, errUnmarshal, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&rr)
		return
	}
	total, err := s.h.GetSearchRecommendationTotalRows(params["query"][0]) // total results
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetchRows, http.StatusInternalServerError)
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
			helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
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
		helper.DecodeError(w, r, s.l, errSearch, http.StatusInternalServerError)
		return
	}
	result := []*model.RecommendationResponse{}
	for _, rr := range search.Data {
		recG, err := s.h.GetRecommendationGenres(rr.ID)
		if err != nil {
			helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
			return
		}
		recK, err := s.h.GetRecommendationKeywords(rr.ID)
		if err != nil {
			helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
			return
		}
		recFinal := &model.RecommendationResponse{
			Recommendation: rr,
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
		helper.DecodeError(w, r, s.l, errMarhsal, http.StatusInternalServerError)
		return
	}
	err = s.rc.Set(rrKey, rr, redisTimeout).Err()
	if err != nil {
		helper.DecodeError(w, r, s.l, errKeySet, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

// SearchUser searches for users.
func (s *Setup) SearchUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.l, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.v.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.h.SearchUser(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.l, errSearch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchGenre searches for genres.
func (s *Setup) SearchGenre(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.l, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.v.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.h.SearchGenre(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.l, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchKeyword searches for keywords.
func (s *Setup) SearchKeyword(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.l, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.v.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.h.SearchKeyword(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.l, errSearch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchSource searches for sources.
func (s *Setup) SearchSource(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.l, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.v.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.h.SearchSource(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.l, errSearch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}
