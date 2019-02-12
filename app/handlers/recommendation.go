package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/gorilla/mux"
)

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	//Redis check start
	var rrKey string
	if params["page"] != nil {
		rrKey = fmt.Sprintf("recommendation?page=%s", params["page"][0])
	} else {
		rrKey = "recommendation"
	}

	val, _ := redisClient.Get(rrKey).Result()

	if val != "" {
		rr := &model.RecommendationPagination{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, "Could not unmarshal the payload", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&rr)
		return
	}
	// Redis check end

	// Start pagination
	total, err := db.GetRecommendationTotalRows() // total results

	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendations total rows", http.StatusInternalServerError)
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
				helper.DecodeError(w, "Could not parse page param to float", http.StatusInternalServerError)
				return
			}

			if page > currentPage {
				currentPage = page
				offset = (currentPage - 1) * limit
			}
		}
	}
	// End pagination

	rec, err := db.GetRecommendations(offset, limit)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendations", http.StatusInternalServerError)
		return
	}

	result := []*model.ResponseRecommendation{}

	for _, r := range rec.Data {
		recG, err := db.GetRecommendationGenres(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the recommendations genres", http.StatusInternalServerError)
			return
		}
		recK, err := db.GetRecommendationKeywords(r.ID)
		if err != nil {
			helper.DecodeError(w, "Could not fetch the recommendations", http.StatusInternalServerError)
			return
		}
		recFinal := &model.ResponseRecommendation{
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
		helper.DecodeError(w, "Could not marshal the payload", http.StatusInternalServerError)
		return
	}
	if err := redisClient.Set(rrKey, rr, redisTimeout).Err(); err != nil {
		helper.DecodeError(w, "Could not set the key", http.StatusInternalServerError)
		return
	}
	// Redis set check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultFinal)
}

func getRecommendation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}

	//Redis check start
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := redisClient.Get(rrKey).Result()

	if val != "" {
		rr := &model.ResponseRecommendation{}
		if err := helper.UnmarshalBinary([]byte(val), rr); err != nil {
			helper.DecodeError(w, "Could not unmarshal the payload", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rr)
		return
	}
	// Redis check end

	rec, err := db.GetRecommendation(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation", http.StatusInternalServerError)
		return
	}

	recG, err := db.GetRecommendationGenres(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation genres", http.StatusInternalServerError)
		return
	}

	recK, err := db.GetRecommendationKeywords(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation keywords", http.StatusInternalServerError)
		return
	}

	response := &model.ResponseRecommendation{
		Recommendation: rec,
		Genres:         recG,
		Keywords:       recK,
	}

	// Redis set check start
	rr, err := helper.MarshalBinary(response)
	if err != nil {
		helper.DecodeError(w, "Could not marshal the payload", http.StatusInternalServerError)
		return
	}

	if err := redisClient.Set(rrKey, rr, redisTimeout).Err(); err != nil {
		helper.DecodeError(w, "Could not set the key", http.StatusInternalServerError)
		return
	}
	// Redis set check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createRecommendation(w http.ResponseWriter, r *http.Request) {

	reqRec := &model.RecommendationCreate{}
	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}

	validate = validator.New()
	if err := validate.Struct(reqRec); err != nil {
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

	rec, err := db.CreateRecommendation(newRec)

	if err != nil {
		helper.DecodeError(w, "Could not create the recommendation", http.StatusInternalServerError)
		return
	}

	// Attaching keywords
	keywords := make(map[int64][]int)
	keywords[rec.ID] = reqRec.Keywords
	_, err = helper.Attach(keywords, "keyword_recommendation", db.DB)
	if err != nil {
		helper.DecodeError(w, "Could not attach the recommendation keywords", http.StatusInternalServerError)
		return
	}

	// Attaching genres
	genres := make(map[int64][]int)
	genres[rec.ID] = reqRec.Genres
	_, err = helper.Attach(genres, "genre_recommendation", db.DB)
	if err != nil {
		helper.DecodeError(w, "Could not attach the recommendation genres", http.StatusInternalServerError)
		return
	}

	// Redis check start
	val, _ := redisClient.Get("recommendation").Result()
	if val != "" {
		_, err = redisClient.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rec)
}

func updateRecommendation(w http.ResponseWriter, r *http.Request) {

	reqRec := &model.RecommendationCreate{}

	if err := json.NewDecoder(r.Body).Decode(reqRec); err != nil {
		helper.DecodeError(w, "Could not decode the body response", http.StatusInternalServerError)
		return
	}

	validate = validator.New()
	if err := validate.Struct(reqRec); err != nil {
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

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}

	// Empty recommendation check
	itemCount, err := db.GetRecommendationItemsTotalRows(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the recommendation items total rows", http.StatusInternalServerError)
		return
	}

	if itemCount == 0 && reqRec.Status == 1 {
		helper.DecodeError(w, "The recommendation is empty or does not exist", http.StatusUnprocessableEntity)
		return
	}

	rec, err := db.UpdateRecommendation(id, &upRec)
	if err != nil {
		helper.DecodeError(w, "Could not update the recommendation", http.StatusInternalServerError)
		return
	}

	// Syncing keywords
	keywords := make(map[int64][]int)
	keywords[rec.ID] = reqRec.Keywords
	_, err = helper.Sync(keywords, "keyword_recommendation", "recommendation_id", db.DB)
	if err != nil {
		helper.DecodeError(w, "Could not sync the recommendation keywords", http.StatusInternalServerError)
		return
	}

	// Syncing genres
	genres := make(map[int64][]int)
	genres[rec.ID] = reqRec.Genres
	_, err = helper.Sync(genres, "genre_recommendation", "recommendation_id", db.DB)
	if err != nil {
		helper.DecodeError(w, "Could not sync the recommendation genres", http.StatusInternalServerError)
		return
	}

	// Redis check start
	val, _ := redisClient.Get("recommendation").Result()
	if val != "" {
		_, err = redisClient.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}

	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ = redisClient.Get(rrKey).Result()
	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&rec)
}

func deleteRecommendation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}

	// Redis check start
	rrKey := fmt.Sprintf("recommendation-%d", id)
	val, _ := redisClient.Get(rrKey).Result()

	if val != "" {
		_, err = redisClient.Unlink(rrKey).Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}

	val, _ = redisClient.Get("recommendation").Result()

	if val != "" {
		_, err = redisClient.Unlink("recommendation").Result()
		if err != nil {
			helper.DecodeError(w, "Could not unlink the key", http.StatusInternalServerError)
			return
		}
	}
	// Redis check end

	if err := db.DeleteRecommendation(id); err != nil {
		helper.DecodeError(w, "Could not delete the recommendation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Recommendation deleted successfully!"})
}
