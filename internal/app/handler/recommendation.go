package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/cyruzin/tome"
	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

// GetRecommendations gets all recommendations.
func (s *Setup) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	var redisKey string

	if params["page"] != nil {
		redisKey = fmt.Sprintf("recommendation?page=%s", params["page"][0])
	} else {
		redisKey = "recommendation"
	}

	var recommendation *model.RecommendationResult

	cache, err := s.CheckCache(redisKey, recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUnmarshal, http.StatusInternalServerError)
		return
	}

	if cache {
		json.NewEncoder(w).Encode(recommendation)
		return
	}

	total, err := s.h.GetRecommendationTotalRows()
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	newPage, err := helper.PageParser(params)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	chapter := &tome.Chapter{NewPage: newPage, TotalResults: total}

	if err := chapter.Paginate(); err != nil {
		helper.DecodeError(w, r, s.l, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := s.h.GetRecommendations(chapter.Offset, chapter.Limit)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	resultFinal := &model.RecommendationResult{
		Data:    result,
		Chapter: chapter,
	}

	err = s.SetCache(redisKey, resultFinal)
	if err != nil {
		helper.DecodeError(w, r, s.l, errMarhsal, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

// GetRecommendation gets a recommendation by ID.
func (s *Setup) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	//Redis check start
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := s.rc.Get(rrKey).Result()

	if val != "" {
		rr := &model.Recommendation{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, r, s.l, errUnmarshal, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rr)
		return
	}
	// Redis check end

	recommentation, err := s.h.GetRecommendation(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	// Redis set check start
	rr, err := helper.MarshalBinary(recommentation)
	if err != nil {
		helper.DecodeError(w, r, s.l, errMarhsal, http.StatusInternalServerError)
		return
	}

	if err := s.rc.Set(rrKey, rr, redisTimeout).Err(); err != nil {
		helper.DecodeError(w, r, s.l, errKeySet, http.StatusInternalServerError)
		return
	}
	// Redis set check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recommentation)
}

// CreateRecommendation creates a new recommendation.
func (s *Setup) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	recommendation := &model.Recommendation{
		UserID:    int64(reqRec.UserID),
		Title:     reqRec.Title,
		Type:      reqRec.Type,
		Body:      reqRec.Body,
		Poster:    reqRec.Poster,
		Backdrop:  reqRec.Backdrop,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	recommendationID, err := s.h.CreateRecommendation(recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.l, errCreate, http.StatusInternalServerError)
		return
	}

	// Attaching keywords
	keywords := make(map[int64][]int)
	keywords[recommendationID] = reqRec.Keywords
	err = s.h.Attach(keywords, "keyword_recommendation")
	if err != nil {
		helper.DecodeError(w, r, s.l, errAttach, http.StatusInternalServerError)
		return
	}

	// Attaching genres
	genres := make(map[int64][]int)
	genres[recommendationID] = reqRec.Genres
	err = s.h.Attach(genres, "genre_recommendation")
	if err != nil {
		helper.DecodeError(w, r, s.l, errAttach, http.StatusInternalServerError)
		return
	}

	// Redis check start
	val, _ := s.rc.Get("recommendation").Result()
	if val != "" {
		_, err = s.rc.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation created successfully!"})
}

// UpdateRecommendation updates a recommendation.
func (s *Setup) UpdateRecommendation(w http.ResponseWriter, r *http.Request) {
	reqRec := &model.RecommendationCreate{}

	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqRec); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	recommendation := model.Recommendation{
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
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	// Empty recommendation check
	itemCount, err := s.h.GetRecommendationItemsTotalRows(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetchRows, http.StatusInternalServerError)
		return
	}

	if itemCount == 0 && reqRec.Status == 1 {
		helper.DecodeError(w, r, s.l, errEmptyRec, http.StatusUnprocessableEntity)
		return
	}

	err = s.h.UpdateRecommendation(id, &recommendation)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUpdate, http.StatusInternalServerError)
		return
	}

	// Syncing keywords
	keywords := make(map[int64][]int)
	keywords[id] = reqRec.Keywords
	err = s.h.Sync(keywords, "keyword_recommendation", "recommendation_id")
	if err != nil {
		helper.DecodeError(w, r, s.l, errSync, http.StatusInternalServerError)
		return
	}

	// Syncing genres
	genres := make(map[int64][]int)
	genres[id] = reqRec.Genres
	err = s.h.Sync(genres, "genre_recommendation", "recommendation_id")
	if err != nil {
		helper.DecodeError(w, r, s.l, errSync, http.StatusInternalServerError)
		return
	}

	// Redis check start
	val, _ := s.rc.Get("recommendation").Result()
	if val != "" {
		_, err = s.rc.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}

	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ = s.rc.Get(rrKey).Result()
	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation updated successfully!"})
}

// DeleteRecommendation deletes a recommendation.
func (s *Setup) DeleteRecommendation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := s.rc.Get(rrKey).Result()

	if val != "" {
		_, err = s.rc.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}

	val, _ = s.rc.Get("recommendation").Result()

	if val != "" {
		_, err = s.rc.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, r, s.l, errKeyUnlink, http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	if err := s.h.DeleteRecommendation(id); err != nil {
		helper.DecodeError(w, r, s.l, errDelete, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation deleted successfully!"})
}

// GetRecommendationsAdmin retrieves the last 10 recommendations
// without filter.
func (s *Setup) GetRecommendationsAdmin(w http.ResponseWriter, r *http.Request) {
	// rec, err := s.h.GetRecommendationsAdmin()
	// if err != nil {
	// 	helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
	// 	return
	// }

	// result := []*model.RecommendationResponse{}

	// for _, rr := range rec.Data {
	// 	recG, err := s.h.GetRecommendationGenres(rr.ID)
	// 	if err != nil {
	// 		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
	// 		return
	// 	}
	// 	recK, err := s.h.GetRecommendationKeywords(rr.ID)
	// 	if err != nil {
	// 		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
	// 		return
	// 	}
	// 	recFinal := &model.RecommendationResponse{
	// 		Recommendation: rr,
	// 		Genres:         recG,
	// 		Keywords:       recK,
	// 	}
	// 	result = append(result, recFinal)
	// }

	// resultFinal := &model.RecommendationPagination{Data: result}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(resultFinal)
}
