package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

// GetRecommendations gets all recommendations.
func (s *Setup) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	//Redis check start
	var rrKey string
	if params["page"] != nil {
		rrKey = fmt.Sprintf("recommendation?page=%s", params["page"][0])
	} else {
		rrKey = "recommendation"
	}

	val, _ := s.rc.Get(rrKey).Result()

	if val != "" {
		rr := &model.RecommendationPagination{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, errUnmarshal, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&rr)
		return
	}
	// Redis check end

	// Start pagination
	total, err := s.h.GetRecommendationTotalRows() // total results

	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}

	var (
		limit       float64 = 10                       // limit per page
		offset      float64                            // offset record
		currentPage float64 = 1                        // current page
		lastPage            = math.Ceil(total / limit) // last page
	)

	// checking if request contains the "page" parameter
	if len(params) > 0 {
		if params["page"][0] != "" {
			page, err := strconv.ParseFloat(params["page"][0], 64)

			if err != nil {
				helper.DecodeError(w, errParseFloat, http.StatusInternalServerError)
				return
			}

			if page > currentPage {
				currentPage = page
				offset = (currentPage - 1) * limit
			}
		}
	}
	// End pagination

	rec, err := s.h.GetRecommendations(offset, limit)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}

	result := []*model.RecommendationResponse{}

	for _, r := range rec.Data {
		recG, err := s.h.GetRecommendationGenres(r.ID)
		if err != nil {
			helper.DecodeError(w, errFetch, http.StatusInternalServerError)
			return
		}
		recK, err := s.h.GetRecommendationKeywords(r.ID)
		if err != nil {
			helper.DecodeError(w, errFetch, http.StatusInternalServerError)
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

	// Redis set check start
	rr, err := helper.MarshalBinary(resultFinal)
	if err != nil {
		helper.DecodeError(w, errMarhsal, http.StatusInternalServerError)
		return
	}
	if err := s.rc.Set(rrKey, rr, redisTimeout).Err(); err != nil {
		helper.DecodeError(w, errKeySet, http.StatusInternalServerError)
		return
	}
	// Redis set check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

// GetRecommendation gets a recommendation by ID.
func (s *Setup) GetRecommendation(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}

	//Redis check start
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := s.rc.Get(rrKey).Result()

	if val != "" {
		rr := &model.RecommendationResponse{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, errUnmarshal, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rr)
		return
	}
	// Redis check end

	rec, err := s.h.GetRecommendation(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}

	recG, err := s.h.GetRecommendationGenres(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}

	recK, err := s.h.GetRecommendationKeywords(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}

	response := &model.RecommendationResponse{
		Recommendation: rec,
		Genres:         recG,
		Keywords:       recK,
	}

	// Redis set check start
	rr, err := helper.MarshalBinary(response)
	if err != nil {
		helper.DecodeError(w, errMarhsal, http.StatusInternalServerError)
		return
	}

	if err := s.rc.Set(rrKey, rr, redisTimeout).Err(); err != nil {
		helper.DecodeError(w, errKeySet, http.StatusInternalServerError)
		return
	}
	// Redis set check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// CreateRecommendation creates a new recommendation.
func (s *Setup) CreateRecommendation(w http.ResponseWriter, r *http.Request) {

	reqRec := &model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	newRec := &model.Recommendation{
		UserID:    int64(reqRec.UserID),
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rec, err := s.h.CreateRecommendation(newRec)

	if err != nil {
		helper.DecodeError(w, errCreate, http.StatusInternalServerError)
		return
	}

	// Attaching keywords
	keywords := make(map[int64][]int)
	keywords[rec.ID] = reqRec.Keywords
	err = s.h.Attach(keywords, "keyword_recommendation")
	if err != nil {
		helper.DecodeError(w, errAttach, http.StatusInternalServerError)
		return
	}

	// Attaching genres
	genres := make(map[int64][]int)
	genres[rec.ID] = reqRec.Genres
	err = s.h.Attach(genres, "genre_recommendation")
	if err != nil {
		helper.DecodeError(w, errAttach, http.StatusInternalServerError)
		return
	}

	// Redis check start
	val, _ := s.rc.Get("recommendation").Result()
	if val != "" {
		_, err = s.rc.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rec)
}

// UpdateRecommendation updates a recommendation.
func (s *Setup) UpdateRecommendation(w http.ResponseWriter, r *http.Request) {

	reqRec := &model.RecommendationCreate{}

	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	upRec := model.Recommendation{
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		Status:    reqRec.Status,
		UpdatedAt: time.Now(),
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}

	// Empty recommendation check
	itemCount, err := s.h.GetRecommendationItemsTotalRows(id)
	if err != nil {
		helper.DecodeError(w, errFetchRows, http.StatusInternalServerError)
		return
	}

	if itemCount == 0 && reqRec.Status == 1 {
		helper.DecodeError(w, errEmptyRec, http.StatusUnprocessableEntity)
		return
	}

	rec, err := s.h.UpdateRecommendation(id, &upRec)
	if err != nil {
		helper.DecodeError(w, errUpdate, http.StatusInternalServerError)
		return
	}

	// Syncing keywords
	keywords := make(map[int64][]int)
	keywords[rec.ID] = reqRec.Keywords
	err = s.h.Sync(keywords, "keyword_recommendation", "recommendation_id")
	if err != nil {
		helper.DecodeError(w, errSync, http.StatusInternalServerError)
		return
	}

	// Syncing genres
	genres := make(map[int64][]int)
	genres[rec.ID] = reqRec.Genres
	err = s.h.Sync(genres, "genre_recommendation", "recommendation_id")
	if err != nil {
		helper.DecodeError(w, errSync, http.StatusInternalServerError)
		return
	}

	// Redis check start
	val, _ := s.rc.Get("recommendation").Result()
	if val != "" {
		_, err = s.rc.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}

	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ = s.rc.Get(rrKey).Result()
	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&rec)
}

// DeleteRecommendation deletes a recommendation.
func (s *Setup) DeleteRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := s.rc.Get(rrKey).Result()

	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}

	val, _ = s.rc.Get("recommendation").Result()

	if val != "" {
		_, err = s.rc.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	if err := s.h.DeleteRecommendation(id); err != nil {
		helper.DecodeError(w, errDelete, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation deleted successfully!"})
}
