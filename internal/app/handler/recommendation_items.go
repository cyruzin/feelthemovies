package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// GetRecommendationItems gets all recommendation items.
func (s *Setup) GetRecommendationItems(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	redisKey := fmt.Sprintf("recommendation_items-%d", id)

	recommendationItemCache := model.RecommendationItemResult{}

	cache, err := s.CheckCache(redisKey, &recommendationItemCache)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		s.ToJSON(w, http.StatusOK, &recommendationItemCache)
		return
	}

	recommendationItems, err := s.model.GetRecommendationItems(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	err = s.SetCache(redisKey, &recommendationItems)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUnmarshal, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationItems)
}

// GetRecommendationItem gets a recommendation item by ID.
func (s *Setup) GetRecommendationItem(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationItem, err := s.model.GetRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &recommendationItem)
}

// CreateRecommendationItem creates a new recommendation item.
func (s *Setup) CreateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	request := model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	yearParsed, err := time.Parse("2006-01-02", request.Year)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseDate, http.StatusInternalServerError)
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

	recommendationID, err := s.model.CreateRecommendationItem(&recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	sources := make(map[int64][]int)
	sources[recommendationID] = request.Sources
	err = s.model.Attach(sources, "recommendation_item_source")
	if err != nil {
		helper.DecodeError(w, r, s.logger, errAttach, http.StatusInternalServerError)
		return
	}

	err = s.RemoveCache(fmt.Sprintf("recommendation_items-%d", recommendationID))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(
		w,
		http.StatusCreated,
		&helper.APIMessage{Message: "Recommendation item created successfully!"},
	)
}

// UpdateRecommendationItem updates a recommendation item.
func (s *Setup) UpdateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseDate, http.StatusInternalServerError)
		return
	}

	recommendation := model.RecommendationItem{
		Name:       reqRec.Name,
		TMDBID:     reqRec.TMDBID,
		Year:       yearParsed,
		Overview:   reqRec.Overview,
		Poster:     reqRec.Poster,
		Backdrop:   reqRec.Backdrop,
		Trailer:    reqRec.Trailer,
		Commentary: reqRec.Commentary,
		MediaType:  reqRec.MediaType,
		UpdatedAt:  time.Now(),
	}

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationItemID, err := s.model.UpdateRecommendationItem(id, &recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	sources := make(map[int64][]int)
	sources[recommendationItemID] = reqRec.Sources
	err = s.model.Sync(sources, "recommendation_item_source", "recommendation_item_id")
	if err != nil {
		helper.DecodeError(w, r, s.logger, errSync, http.StatusInternalServerError)
		return
	}

	err = s.RemoveCache(fmt.Sprintf("recommendation_items-%d", recommendationItemID))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	s.ToJSON(
		w,
		http.StatusOK,
		&helper.APIMessage{Message: "Recommendation item updated successfully!"},
	)
}

// DeleteRecommendationItem deletes a recommendation item.
func (s *Setup) DeleteRecommendationItem(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	recommendationItem, err := s.model.GetRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	err = s.RemoveCache(
		fmt.Sprintf("recommendation_items-%d", recommendationItem.RecommendationID),
	)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errKeyUnlink, http.StatusInternalServerError)
		return
	}

	err = s.model.DeleteRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(
		w,
		http.StatusOK,
		&helper.APIMessage{Message: "Recommendation item deleted successfully!"},
	)
}
