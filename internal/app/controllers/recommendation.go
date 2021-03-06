package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
	"github.com/cyruzin/tome"
	"github.com/go-chi/chi"

	model "github.com/cyruzin/feelthemovies/internal/app/models"
)

// GetRecommendations gets all recommendations.
func (s *Setup) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	redisKey := s.GenerateCacheKey(params, "recommendation")

	recommendationsCache := model.RecommendationResult{}

	ctx := r.Context()

	cache, err := s.CheckCache(ctx, redisKey, &recommendationsCache)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationsCache)
		return
	}

	total, err := s.model.GetRecommendationTotalRows(ctx)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
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

	result, err := s.model.GetRecommendations(ctx, chapter.Offset, chapter.Limit)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	recommendations := model.RecommendationResult{Data: result, Chapter: &chapter}

	err = s.SetCache(ctx, redisKey, recommendations)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errMarshal, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendations)
}

// GetRecommendation gets a recommendation by ID.
func (s *Setup) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	redisKey := s.GenerateCacheKey(nil, fmt.Sprintf("recommendation-%d", id))

	recommendationCache := model.Recommendation{}

	ctx := r.Context()

	cache, err := s.CheckCache(ctx, redisKey, &recommendationCache)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationCache)
		return
	}

	recommendation, err := s.model.GetRecommendation(ctx, id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.SetCache(ctx, redisKey, &recommendation)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errMarshal, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendation)
}

// CreateRecommendation creates a new recommendation.
func (s *Setup) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	request := model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	recommendation := model.Recommendation{
		UserID:    request.UserID,
		Title:     request.Title,
		Type:      request.Type,
		Body:      request.Body,
		Poster:    request.Poster,
		Backdrop:  request.Backdrop,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	recommendationID, err := s.model.CreateRecommendation(ctx, &recommendation)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	keywords := make(map[int64][]int)
	keywords[recommendationID] = request.Keywords
	err = s.model.Attach(ctx, keywords, "keyword_recommendation")
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errAttach, http.StatusInternalServerError)
		return
	}

	genres := make(map[int64][]int)
	genres[recommendationID] = request.Genres
	err = s.model.Attach(ctx, genres, "genre_recommendation")
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errAttach, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache(ctx, "recommendation"); err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &errhandler.APIMessage{Message: "Recommendation created successfully!"})
}

// UpdateRecommendation updates a recommendation.
func (s *Setup) UpdateRecommendation(w http.ResponseWriter, r *http.Request) {
	request := model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	recommendation := model.Recommendation{
		Title:     request.Title,
		Type:      request.Type,
		Body:      request.Body,
		Poster:    request.Poster,
		Backdrop:  request.Backdrop,
		Status:    request.Status,
		UpdatedAt: time.Now(),
	}

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	itemCount, err := s.model.GetRecommendationItemsTotalRows(ctx, id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetchRows, http.StatusInternalServerError)
		return
	}

	if itemCount == 0 && request.Status == 1 {
		errhandler.DecodeError(w, r, s.logger, errEmptyRec, http.StatusUnprocessableEntity)
		return
	}

	err = s.model.UpdateRecommendation(ctx, id, &recommendation)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	keywords := make(map[int64][]int)
	keywords[id] = request.Keywords
	err = s.model.Sync(ctx, keywords, "keyword_recommendation", "recommendation_id")
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSync, http.StatusInternalServerError)
		return
	}

	genres := make(map[int64][]int)
	genres[id] = request.Genres
	err = s.model.Sync(ctx, genres, "genre_recommendation", "recommendation_id")
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSync, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache(ctx, "recommendation"); err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache(ctx, fmt.Sprintf("recommendation-%d", id)); err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "Recommendation updated successfully!"})
}

// DeleteRecommendation deletes a recommendation.
func (s *Setup) DeleteRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.RemoveCache(ctx, fmt.Sprintf("recommendation-%d", id)); err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache(ctx, "recommendation"); err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteRecommendation(ctx, id); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "Recommendation deleted successfully!"})
}

// GetRecommendationsAdmin get the latest recommendations
// without status filter.
func (s *Setup) GetRecommendationsAdmin(w http.ResponseWriter, r *http.Request) {
	recommendations, err := s.model.GetRecommendationsAdmin(r.Context())
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendations)
}

// GetRecommendationGenres gets all genres
// of a specific recommendation.
func (s *Setup) GetRecommendationGenres(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationGenres, err := s.model.GetRecommendationGenres(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationGenres)
}

// GetRecommendationKeywords gets all keywords
// of a specific recommendation.
func (s *Setup) GetRecommendationKeywords(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationKeywords, err := s.model.GetRecommendationKeywords(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationKeywords)
}
