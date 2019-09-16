package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// GetRecommendationItems gets all recommendation items.
func (s *Setup) GetRecommendationItems(w http.ResponseWriter, r *http.Request) {
	id, err := helper.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	//Redis check start
	redisKey := fmt.Sprintf("recommendation_items-%d", id)

	recommendationItem := &model.RecommendationItemFinal{}

	cache, err := s.CheckCache(redisKey, recommendationItem)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(recommendationItem)
		return
	}
	// Redis check end

	rec, err := s.h.GetRecommendationItems(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	result := []*model.RecommendationItemResponse{}
	for _, rr := range rec.Data {
		recS, err := s.h.GetRecommendationItemSources(rr.ID)
		if err != nil {
			helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
			return
		}

		recFinal := &model.RecommendationItemResponse{
			RecommendationItem: rr,
			Sources:            recS,
		}

		result = append(result, recFinal)
	}

	resultFinal := &model.RecommendationItemFinal{Data: result}

	// Redis set start
	err = s.SetCache(redisKey, resultFinal)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUnmarshal, http.StatusInternalServerError)
		return
	}
	// Redis set end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

// GetRecommendationItem gets a recommendation item by ID.
func (s *Setup) GetRecommendationItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	rec, err := s.h.GetRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	recS, err := s.h.GetRecommendationItemSources(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	response := &model.RecommendationItemResponse{
		RecommendationItem: rec,
		Sources:            recS,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// CreateRecommendationItem creates a new recommendation item.
func (s *Setup) CreateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	// Parsing string to time.Time
	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseDate, http.StatusInternalServerError)
		return
	}

	newRec := &model.RecommendationItem{
		RecommendationID: reqRec.RecommendationID,
		Name:             reqRec.Name,
		TMDBID:           reqRec.TMDBID,
		Year:             yearParsed,
		Overview:         reqRec.Overview,
		Poster:           reqRec.Poster,
		Backdrop:         reqRec.Backdrop,
		Trailer:          reqRec.Trailer,
		Commentary:       reqRec.Commentary,
		MediaType:        reqRec.MediaType,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	rec, err := s.h.CreateRecommendationItem(newRec)
	if err != nil {
		helper.DecodeError(w, r, s.l, errCreate, http.StatusInternalServerError)
		return
	}

	// Attaching sources
	sources := make(map[int64][]int)
	sources[rec.ID] = reqRec.Sources
	err = s.h.Attach(sources, "recommendation_item_source")
	if err != nil {
		helper.DecodeError(w, r, s.l, errAttach, http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, _ := s.rc.Get(rrKey).Result()
	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&rec)

}

// UpdateRecommendationItem updates a recommendation item.
func (s *Setup) UpdateRecommendationItem(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationItemCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	yearParsed, err := time.Parse("2006-01-02", reqRec.Year)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseDate, http.StatusInternalServerError)
		return
	}

	upRec := &model.RecommendationItem{
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

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	rec, err := s.h.UpdateRecommendationItem(id, upRec)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUpdate, http.StatusInternalServerError)
		return
	}

	// Syncing sources
	sources := make(map[int64][]int)
	sources[rec.ID] = reqRec.Sources
	err = s.h.Sync(sources, "recommendation_item_source", "recommendation_item_id")
	if err != nil {
		helper.DecodeError(w, r, s.l, errSync, http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, _ := s.rc.Get(rrKey).Result()
	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&rec)
}

// DeleteRecommendationItem ...
func (s *Setup) DeleteRecommendationItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	rec, err := s.h.GetRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation_items-%d", rec.RecommendationID)
	val, _ := s.rc.Get(rrKey).Result()
	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}

	err = s.h.DeleteRecommendationItem(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errDelete, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation item deleted successfully!"})
}
