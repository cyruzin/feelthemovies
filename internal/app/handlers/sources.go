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

// GetSources ...
func (s *Setup) GetSources(w http.ResponseWriter, r *http.Request) {
	so, err := s.h.GetSources()
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&so)
}

// GetSource ...
func (s *Setup) GetSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	so, err := s.h.GetSource(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&so)
}

// CreateSource ...
func (s *Setup) CreateSource(w http.ResponseWriter, r *http.Request) {
	reqS := &model.Source{}
	if err := json.NewDecoder(r.Body).Decode(reqS); err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
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
		helper.DecodeError(w, errCreate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&so)
}

// UpdateSource ...
func (s *Setup) UpdateSource(w http.ResponseWriter, r *http.Request) {
	reqS := &model.Source{}
	if err := json.NewDecoder(r.Body).Decode(reqS); err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
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
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	so, err := s.h.UpdateSource(id, &upS)
	if err != nil {
		helper.DecodeError(w, errUpdate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&so)
}

// DeleteSource ...
func (s *Setup) DeleteSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	if err := s.h.DeleteSource(id); err != nil {
		helper.DecodeError(w, errDelete, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Source deleted successfully!"})
}
