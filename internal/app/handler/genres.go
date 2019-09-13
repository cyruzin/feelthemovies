package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

// GetGenres gets all genres.
func (s *Setup) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := s.h.GetGenres(20)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&genres)
}

// GetGenre gets a genre by ID.
func (s *Setup) GetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	genre, err := s.h.GetGenre(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&genre)
}

// CreateGenre creates a new genre.
func (s *Setup) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.Genre

	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(genre); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	genre.CreatedAt = time.Now()
	genre.UpdatedAt = time.Now()

	err = s.h.CreateGenre(&genre)
	if err != nil {
		helper.DecodeError(w, r, s.l, errCreate, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Genre created successfully!"})
}

// UpdateGenre updates a genre.
func (s *Setup) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.Genre

	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(genre); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	genre.UpdatedAt = time.Now()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.h.UpdateGenre(id, &genre)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUpdate, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Genre updated successfully!"})
}

// DeleteGenre deletes a genre.
func (s *Setup) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.h.DeleteGenre(id); err != nil {
		helper.DecodeError(w, r, s.l, errDelete, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Genre deleted successfully!"})
}
