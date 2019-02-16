package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

// GetGenres ...
func (s *Setup) GetGenres(w http.ResponseWriter, r *http.Request) {
	g, err := s.h.GetGenres()
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&g)
}

// GetGenre ...
func (s *Setup) GetGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	g, err := s.h.GetGenre(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&g)
}

// CreateGenre ...
func (s *Setup) CreateGenre(w http.ResponseWriter, r *http.Request) {
	reqG := &model.Genre{}
	err := json.NewDecoder(r.Body).Decode(reqG)
	if err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqG); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	newG := model.Genre{
		Name:      reqG.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	g, err := s.h.CreateGenre(&newG)
	if err != nil {
		helper.DecodeError(w, errCreate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&g)
}

// UpdateGenre ...
func (s *Setup) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	reqG := &model.Genre{}
	err := json.NewDecoder(r.Body).Decode(reqG)
	if err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqG); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	upG := model.Genre{
		Name:      reqG.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	g, err := s.h.UpdateGenre(id, &upG)
	if err != nil {
		helper.DecodeError(w, errUpdate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&g)
}

// DeleteGenre ...
func (s *Setup) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	if err := s.h.DeleteGenre(id); err != nil {
		helper.DecodeError(w, errDelete, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Genre deleted successfully!"})
}
