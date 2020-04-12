package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
)

// GetRecommendationItems gets all recommendation items.
func (s *Setup) GetRecommendationItems(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	redisKey := fmt.Sprintf("recommendation_items-%d", id)

	recommendationItemCache := model.RecommendationItemResult{}

	cache, err := s.CheckCache(ctx, redisKey, &recommendationItemCache)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationItemCache)
		return
	}

	recommendationItems, err := s.model.GetRecommendationItems(ctx, id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	err = s.SetCache(ctx, redisKey, &recommendationItems)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationItems)
}

// GetRecommendationItem gets a recommendation item by ID.
func (s *Setup) GetRecommendationItem(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationItem, err := s.model.GetRecommendationItem(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationItem)
}

// CreateRecommendationItem creates a new recommendation item.
func (s *Setup) CreateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	request := model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	yearParsed, err := time.Parse("2006-01-02", request.Year)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseDate, http.StatusInternalServerError)
		return
	}

	recommendation := model.RecommendationItem{
		RecommendationID: request.RecommendationID,
		Name:             request.Name,
		TMDBID:           request.TMDBID,
		Year:             yearParsed,
		Overview:         request.Overview,
		Poster:           request.Poster,
		Backdrop:         request.Backdrop,
		Trailer:          request.Trailer,
		Commentary:       request.Commentary,
		MediaType:        request.MediaType,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	recommendationID, err := s.model.CreateRecommendationItem(ctx, &recommendation)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	sources := make(map[int64][]int)
	sources[recommendationID] = request.Sources
	err = s.model.Attach(ctx, sources, "recommendation_item_source")
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errAttach, http.StatusInternalServerError)
		return
	}

	err = s.RemoveCache(ctx, fmt.Sprintf("recommendation_items-%d", request.RecommendationID))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(
		w,
		http.StatusCreated,
		&errhandler.APIMessage{Message: "Recommendation item created successfully!"},
	)
}

// UpdateRecommendationItem updates a recommendation item.
func (s *Setup) UpdateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	request := &model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	yearParsed, err := time.Parse("2006-01-02", request.Year)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseDate, http.StatusInternalServerError)
		return
	}

	recommendation := model.RecommendationItem{
		Name:       request.Name,
		TMDBID:     request.TMDBID,
		Year:       yearParsed,
		Overview:   request.Overview,
		Poster:     request.Poster,
		Backdrop:   request.Backdrop,
		Trailer:    request.Trailer,
		Commentary: request.Commentary,
		MediaType:  request.MediaType,
		UpdatedAt:  time.Now(),
	}

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.model.UpdateRecommendationItem(ctx, id, &recommendation)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	sources := make(map[int64][]int)
	sources[id] = request.Sources
	err = s.model.Sync(ctx, sources, "recommendation_item_source", "recommendation_item_id")
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errSync, http.StatusInternalServerError)
		return
	}

	err = s.RemoveCache(ctx, fmt.Sprintf("recommendation_items-%d", request.RecommendationID))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(
		w,
		http.StatusOK,
		&errhandler.APIMessage{Message: "Recommendation item updated successfully!"},
	)
}

// DeleteRecommendationItem deletes a recommendation item.
func (s *Setup) DeleteRecommendationItem(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	recommendationItem, err := s.model.GetRecommendationItem(ctx, id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	err = s.RemoveCache(
		ctx,
		fmt.Sprintf("recommendation_items-%d", recommendationItem.RecommendationID),
	)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	err = s.model.DeleteRecommendationItem(ctx, id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(
		w,
		http.StatusOK,
		&errhandler.APIMessage{Message: "Recommendation item deleted successfully!"},
	)
}

// GetRecommendationItemSources gets all sources
// of a specific recommendation item.
func (s *Setup) GetRecommendationItemSources(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationItemSources, err := s.model.GetRecommendationItemSources(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationItemSources)
}
