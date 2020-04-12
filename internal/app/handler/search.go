package handler

import (
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
	"github.com/cyruzin/tome"
)

// SearchRecommendation searches for recommendations.
func (s *Setup) SearchRecommendation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		errhandler.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}

	query := params["query"][0]

	ctx := r.Context()

	if err := s.validator.VarCtx(ctx, query, "required"); err != nil {
		validation.SearchValidatorMessage(w)
		return
	}

	redisKey := s.GenerateCacheKey(params, "")

	recommendationCache := model.RecommendationResult{}

	cache, err := s.CheckCache(ctx, redisKey, &recommendationCache)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationCache)
		return
	}

	total, err := s.model.GetSearchRecommendationTotalRows(ctx, query)

	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetchRows, http.StatusInternalServerError)
		return
	}

	if total == 0 {
		s.ToJSON(w, http.StatusOK, &model.RecommendationResult{})
		return
	}

	newPage, err := s.PageParser(params)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	chapter := tome.Chapter{NewPage: newPage, TotalResults: total}

	if err := chapter.Paginate(); err != nil {
		errhandler.DecodeError(w, r, s.logger, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := s.model.SearchRecommendation(ctx, chapter.Offset, chapter.Limit, query)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	recommendation := model.RecommendationResult{Data: result, Chapter: &chapter}

	err = s.SetCache(ctx, redisKey, &recommendation)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeySet, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendation)
}

// SearchUser searches for users.
func (s *Setup) SearchUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		errhandler.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := s.validator.VarCtx(ctx, params["query"][0], "required"); err != nil {
		validation.SearchValidatorMessage(w)
		return
	}

	search, err := s.model.SearchUser(ctx, params["query"][0])
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &search)
}

// SearchGenre searches for genres.
func (s *Setup) SearchGenre(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		errhandler.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := s.validator.VarCtx(ctx, params["query"][0], "required"); err != nil {
		validation.SearchValidatorMessage(w)
		return
	}

	search, err := s.model.SearchGenre(ctx, params["query"][0])
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &search)
}

// SearchKeyword searches for keywords.
func (s *Setup) SearchKeyword(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		errhandler.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := s.validator.VarCtx(ctx, params["query"][0], "required"); err != nil {
		validation.SearchValidatorMessage(w)
		return
	}

	search, err := s.model.SearchKeyword(ctx, params["query"][0])
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &search)
}

// SearchSource searches for sources.
func (s *Setup) SearchSource(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if len(params) == 0 {
		errhandler.DecodeError(w, r, s.logger, errQueryField, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := s.validator.VarCtx(ctx, params["query"][0], "required"); err != nil {
		validation.SearchValidatorMessage(w)
		return
	}

	search, err := s.model.SearchSource(ctx, params["query"][0])
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSearch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &search)
}
