package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

// GetKeywords gets all keywords.
func (s *Setup) GetKeywords(w http.ResponseWriter, r *http.Request) {
	k, err := s.h.GetKeywords()
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&k)
}

// GetKeyword gets a keyword by ID.
func (s *Setup) GetKeyword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	k, err := s.h.GetKeyword(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&k)
}

// CreateKeyword creates a new keyword.
func (s *Setup) CreateKeyword(w http.ResponseWriter, r *http.Request) {
	reqK := &model.Keyword{}
	err := json.NewDecoder(r.Body).Decode(reqK)
	if err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqK); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	newK := model.Keyword{
		Name:      reqK.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	k, err := s.h.CreateKeyword(&newK)
	if err != nil {
		helper.DecodeError(w, errCreate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&k)
}

// UpdateKeyword updates a keyword.
func (s *Setup) UpdateKeyword(w http.ResponseWriter, r *http.Request) {
	reqK := &model.Keyword{}
	err := json.NewDecoder(r.Body).Decode(reqK)
	if err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqK); err != nil {
		helper.ValidatorMessage(w, err)
	}
	upK := model.Keyword{
		Name:      reqK.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	k, err := s.h.UpdateKeyword(id, &upK)
	if err != nil {
		helper.DecodeError(w, errUpdate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&k)
}

// DeleteKeyword deletes a keyword.
func (s *Setup) DeleteKeyword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	if err := s.h.DeleteKeyword(id); err != nil {
		helper.DecodeError(w, errDelete, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Keyword deleted successfully!"})
}