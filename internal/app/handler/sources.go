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

// GetSources gets all sources.
func (s *Setup) GetSources(w http.ResponseWriter, r *http.Request) {
	so, err := s.h.GetSources()
	if err != nil {
		helper.DecodeError(w,  s.l, err, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&so)
}

// GetSource gets a source by ID.
func (s *Setup) GetSource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errParseInt, http.StatusInternalServerError)
		return
	}
	so, err := s.h.GetSource(id)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&so)
}

// CreateSource creates a new source.
func (s *Setup) CreateSource(w http.ResponseWriter, r *http.Request) {
	reqS := &model.Source{}
	if err := json.NewDecoder(r.Body).Decode(reqS); err != nil {
		helper.DecodeError(w,  s.l, err, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqS); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	newS := model.Source{
		Name:      reqS.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	so, err := s.h.CreateSource(&newS)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errCreate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&so)
}

// UpdateSource updates a source.
func (s *Setup) UpdateSource(w http.ResponseWriter, r *http.Request) {
	reqS := &model.Source{}
	if err := json.NewDecoder(r.Body).Decode(reqS); err != nil {
		helper.DecodeError(w,  s.l, err, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqS); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	upS := model.Source{
		Name:      reqS.Name,
		UpdatedAt: time.Now(),
	}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errParseInt, http.StatusInternalServerError)
		return
	}
	so, err := s.h.UpdateSource(id, &upS)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errUpdate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&so)
}

// DeleteSource deletes a source.
func (s *Setup) DeleteSource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errParseInt, http.StatusInternalServerError)
		return
	}
	if err := s.h.DeleteSource(id); err != nil {
		helper.DecodeError(w,  s.l, err, errDelete, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Source deleted successfully!"})
}
