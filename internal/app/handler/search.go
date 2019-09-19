package handler

import (
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/cyruzin/tome"
)

// SearchRecommendation searches for recommendations.
func (s *Setup) SearchRecommendation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		helper.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}

	query := params["query"][0]

	if err := s.validator.Var(query, "required"); err != nil {
		helper.SearchValidatorMessage(w)
		return
	}

	redisKey := s.GenerateCacheKey(params, "")

	recommendationCache := model.RecommendationResult{}

	cache, err := s.CheckCache(redisKey, &recommendationCache)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationCache)
		return
	}

	total, err := s.model.GetSearchRecommendationTotalRows(query)

	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetchRows, http.StatusInternalServerError)
		return
	}

	if total == 0 {
		s.ToJSON(w, http.StatusOK, &model.RecommendationResult{})
		return
	}

	newPage, err := s.PageParser(params)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	chapter := tome.Chapter{NewPage: newPage, TotalResults: total}

	if err := chapter.Paginate(); err != nil {
		helper.DecodeError(w, r, s.logger, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := s.model.SearchRecommendation(chapter.Offset, chapter.Limit, query)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	recommendation := model.RecommendationResult{Data: result, Chapter: &chapter}

	err = s.SetCache(redisKey, &recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errKeySet, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendation)
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

	s.ToJSON(w, http.StatusOK, &search)
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
		helper.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &search)
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

	s.ToJSON(w, http.StatusOK, &search)
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

	s.ToJSON(w, http.StatusOK, &search)
}
