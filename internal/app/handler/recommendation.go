package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/cyruzin/tome"
	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

// GetRecommendations gets all recommendations.
func (s *Setup) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	redisKey := s.GenerateCacheKey(params, "recommendation")

	recommendationsCache := model.RecommendationResult{}

	cache, err := s.CheckCache(redisKey, &recommendationsCache)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationsCache)
		return
	}

	total, err := s.model.GetRecommendationTotalRows()
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
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

	result, err := s.model.GetRecommendations(chapter.Offset, chapter.Limit)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	recommendations := model.RecommendationResult{Data: result, Chapter: &chapter}

	err = s.SetCache(redisKey, recommendations)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errMarshal, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendations)
}

// GetRecommendation gets a recommendation by ID.
func (s *Setup) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	redisKey := s.GenerateCacheKey(nil, fmt.Sprintf("recommendation-%d", id))

	recommendationCache := model.Recommendation{}

	cache, err := s.CheckCache(redisKey, &recommendationCache)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationCache)
		return
	}

	recommendation, err := s.model.GetRecommendation(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.SetCache(redisKey, &recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errMarshal, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendation)
}

// CreateRecommendation creates a new recommendation.
func (s *Setup) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	request := model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(w, err)
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

	recommendationID, err := s.model.CreateRecommendation(&recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	keywords := make(map[int64][]int)
	keywords[recommendationID] = request.Keywords
	err = s.model.Attach(keywords, "keyword_recommendation")
	if err != nil {
		helper.DecodeError(w, r, s.logger, errAttach, http.StatusInternalServerError)
		return
	}

	genres := make(map[int64][]int)
	genres[recommendationID] = request.Genres
	err = s.model.Attach(genres, "genre_recommendation")
	if err != nil {
		helper.DecodeError(w, r, s.logger, errAttach, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache("recommendation"); err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &helper.APIMessage{Message: "Recommendation created successfully!"})
}

// UpdateRecommendation updates a recommendation.
func (s *Setup) UpdateRecommendation(w http.ResponseWriter, r *http.Request) {
	request := model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(w, err)
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
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	itemCount, err := s.model.GetRecommendationItemsTotalRows(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetchRows, http.StatusInternalServerError)
		return
	}

	if itemCount == 0 && request.Status == 1 {
		helper.DecodeError(w, r, s.logger, errEmptyRec, http.StatusUnprocessableEntity)
		return
	}

	err = s.model.UpdateRecommendation(id, &recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	keywords := make(map[int64][]int)
	keywords[id] = request.Keywords
	err = s.model.Sync(keywords, "keyword_recommendation", "recommendation_id")
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSync, http.StatusInternalServerError)
		return
	}

	genres := make(map[int64][]int)
	genres[id] = request.Genres
	err = s.model.Sync(genres, "genre_recommendation", "recommendation_id")
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSync, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache("recommendation"); err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache(fmt.Sprintf("recommendation-%d", id)); err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &helper.APIMessage{Message: "Recommendation updated successfully!"})
}

// DeleteRecommendation deletes a recommendation.
func (s *Setup) DeleteRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache(fmt.Sprintf("recommendation-%d", id)); err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	if err := s.RemoveCache("recommendation"); err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteRecommendation(id); err != nil {
		helper.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &helper.APIMessage{Message: "Recommendation deleted successfully!"})
}

// GetRecommendationsAdmin ...
func (s *Setup) GetRecommendationsAdmin(w http.ResponseWriter, r *http.Request) {
	recommendations, err := s.model.GetRecommendationsAdmin()
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendations)
}
