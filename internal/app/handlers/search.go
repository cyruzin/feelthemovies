package handlers

import (
	"encoding/json"
	"net/http"

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

	search, err := model.SearchRecommendation(s.Query, s.Type, db)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(search)
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
