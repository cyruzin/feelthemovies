package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	validator "gopkg.in/go-playground/validator.v9"
)

func searchRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var s model.Search

	err = json.NewDecoder(r.Body).Decode(&s)

	validate = validator.New()
	err = validate.Struct(s)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}

	// Start pagination
	params := r.URL.Query()

	total, err := model.GetSearchRecommendationTotalRows(s.Query, db) // total results
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
				log.Println(err)
			}

			if page > currentPage {
				currentPage = page
				offset = (currentPage - 1) * limit
			}
		}
	}

	// End pagination

	search, err := model.SearchRecommendation(offset, limit, s.Query, db)

	result := []*model.ResponseRecommendation{}

	for _, r := range search.Data {
		recG, err := model.GetRecommendationGenres(r.ID, db)
		recK, err := model.GetRecommendationKeywords(r.ID, db)

		if err != nil {
			log.Println(err)
		}

		recFinal := model.ResponseRecommendation{}

		recFinal.Recommendation = r
		recFinal.Genres = recG
		recFinal.Keywords = recK

		result = append(result, &recFinal)
	}

	resultFinal := model.RecommendationPagination{}

	resultFinal.Data = result
	resultFinal.CurrentPage = currentPage
	resultFinal.LastPage = lastPage
	resultFinal.PerPage = limit
	resultFinal.Total = total

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resultFinal)
	}
}

func searchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var s model.Search

	err = json.NewDecoder(r.Body).Decode(&s)

	validate = validator.New()
	err = validate.Struct(s)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}

	search, err := model.SearchUser(s.Query, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}

}

func searchGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var s model.Search

	err = json.NewDecoder(r.Body).Decode(&s)

	validate = validator.New()
	err = validate.Struct(s)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}

	search, err := model.SearchGenre(s.Query, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}

}

func searchKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var s model.Search

	err = json.NewDecoder(r.Body).Decode(&s)

	validate = validator.New()
	err = validate.Struct(s)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}

	search, err := model.SearchKeyword(s.Query, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}

}

func searchSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var s model.Search

	err = json.NewDecoder(r.Body).Decode(&s)

	validate = validator.New()
	err = validate.Struct(s)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("query field is required")
		return
	}

	search, err := model.SearchSource(s.Query, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
	}

}
