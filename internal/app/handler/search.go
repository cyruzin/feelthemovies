package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// SearchRecommendation searches for recommendations.
func (s *Setup) SearchRecommendation(w http.ResponseWriter, r *http.Request) {
	// params := r.URL.Query()
	// if len(params) == 0 {
	// 	helper.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
	// 	return
	// }
	// if err := s.v.Var(params["query"][0], "required"); err != nil {
	// 	helper.SearchValidatorMessage(w)
	// 	return
	// }
	// //Redis check
	// var rrKey string
	// if params["page"] != nil {
	// 	rrKey = fmt.Sprintf(
	// 		"?query=%s?page=%s",
	// 		params["query"][0], params["page"][0],
	// 	)
	// } else {
	// 	rrKey = params["query"][0]
	// }

	// val, _ := s.rc.Get(rrKey).Result()
	// if val != "" {
	// 	rr := &model.RecommendationPagination{}
	// 	if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
	// 		helper.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(&rr)
	// 	return
	// }

	// // Start pagination
	// total, err := s.model.GetSearchRecommendationTotalRows(params["query"][0]) // total results

	// if err != nil {
	// 	helper.DecodeError(w, r, s.logger, errFetchRows, http.StatusInternalServerError)
	// 	return
	// }

	// if total == 0 { // Fix for total rows equal to zero.
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(&model.RecommendationPagination{})
	// 	return
	// }

	// newPage := 1
	// if params["page"] != nil && params["page"][0] != "" {
	// 	newPage, err = strconv.Atoi(params["page"][0])
	// 	if err != nil {
	// 		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	// chapter := &tome.Chapter{
	// 	NewPage:      newPage,
	// 	TotalResults: total,
	// }

	// if err := chapter.Paginate(); err != nil {
	// 	helper.DecodeError(w, r, s.logger, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// // End pagination

	// search, err := s.model.SearchRecommendation(chapter.Offset, chapter.Limit, params["query"][0])
	// if err != nil {
	// 	helper.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
	// 	return
	// }

	// result := []*model.RecommendationResponse{}

	// for _, rr := range search.Data {
	// 	recG, err := s.model.GetRecommendationGenres(rr.ID)
	// 	if err != nil {
	// 		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
	// 		return
	// 	}
	// 	recK, err := s.model.GetRecommendationKeywords(rr.ID)
	// 	if err != nil {
	// 		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
	// 		return
	// 	}
	// 	recFinal := &model.RecommendationResponse{
	// 		Recommendation: rr,
	// 		Genres:         recG,
	// 		Keywords:       recK,
	// 	}
	// 	result = append(result, recFinal)
	// }

	// resultFinal := &model.RecommendationPagination{
	// 	Data:    result,
	// 	Chapter: chapter,
	// }

	// // Redis set
	// rr, err := helper.MarshalBinary(resultFinal)
	// if err != nil {
	// 	helper.DecodeError(w, r, s.logger, errMarhsal, http.StatusInternalServerError)
	// 	return
	// }
	// err = s.rc.Set(rrKey, rr, redisTimeout).Err()
	// if err != nil {
	// 	helper.DecodeError(w, r, s.logger, errKeySet, http.StatusInternalServerError)
	// 	return
	// }
	// // Redis set check end

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(resultFinal)
}

// SearchUser searches for users.
func (s *Setup) SearchUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.validator.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.model.SearchUser(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchGenre searches for genres.
func (s *Setup) SearchGenre(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.validator.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.model.SearchGenre(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.logger, "Could not do the search", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchKeyword searches for keywords.
func (s *Setup) SearchKeyword(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.validator.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.model.SearchKeyword(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}

// SearchSource searches for sources.
func (s *Setup) SearchSource(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}
	if err := s.validator.Var(params["query"][0], "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}
	search, err := s.model.SearchSource(params["query"][0])
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&search)
}
