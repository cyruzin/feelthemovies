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
	genres, err := s.model.GetGenres(20)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &genres)
}

// GetGenre gets a genre by ID.
func (s *Setup) GetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	genre, err := s.model.GetGenre(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &genre)
}

// CreateGenre creates a new genre.
func (s *Setup) CreateGenre(w http.ResponseWriter, r *http.Request) {
	genre := model.Genre{}

	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(genre); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	genre.CreatedAt = time.Now()
	genre.UpdatedAt = time.Now()

	err = s.model.CreateGenre(&genre)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &helper.APIMessage{Message: "Genre created successfully!"})
}

// UpdateGenre updates a genre.
func (s *Setup) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	genre := model.Genre{}

	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(genre); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	genre.UpdatedAt = time.Now()

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.model.UpdateGenre(id, &genre)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &helper.APIMessage{Message: "Genre updated successfully!"})
}

// DeleteGenre deletes a genre.
func (s *Setup) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteGenre(id); err != nil {
		helper.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &helper.APIMessage{Message: "Genre deleted successfully!"})
}
