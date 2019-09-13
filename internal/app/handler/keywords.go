package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

// GetKeywords gets all keywords.
func (s *Setup) GetKeywords(w http.ResponseWriter, r *http.Request) {
	keywords, err := s.h.GetKeywords()
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&keywords)
}

// GetKeyword gets a keyword by ID.
func (s *Setup) GetKeyword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	keyword, err := s.h.GetKeyword(id)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&keyword)
}

// CreateKeyword creates a new keyword.
func (s *Setup) CreateKeyword(w http.ResponseWriter, r *http.Request) {
	var keyword model.Keyword

	err := json.NewDecoder(r.Body).Decode(&keyword)
	if err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(keyword); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	keyword.CreatedAt = time.Now()
	keyword.UpdatedAt = time.Now()

	err = s.h.CreateKeyword(&keyword)
	if err != nil {
		helper.DecodeError(w, r, s.l, errCreate, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Keyword created successfully!"})
}

// UpdateKeyword updates a keyword.
func (s *Setup) UpdateKeyword(w http.ResponseWriter, r *http.Request) {
	var keyword model.Keyword

	err := json.NewDecoder(r.Body).Decode(&keyword)
	if err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(keyword); err != nil {
		helper.ValidatorMessage(w, err)
	}

	keyword.UpdatedAt = time.Now()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.h.UpdateKeyword(id, &keyword)
	if err != nil {
		helper.DecodeError(w, r, s.l, errUpdate, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Keyword updated successfully!"})
}

// DeleteKeyword deletes a keyword.
func (s *Setup) DeleteKeyword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w, r, s.l, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.h.DeleteKeyword(id); err != nil {
		helper.DecodeError(w, r, s.l, errDelete, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Keyword deleted successfully!"})
}
