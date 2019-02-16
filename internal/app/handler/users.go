package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/gorilla/mux"
)

// GetUsers ...
func (s *Setup) GetUsers(w http.ResponseWriter, r *http.Request) {
	u, err := s.h.GetUsers()
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

// GetUser ...
func (s *Setup) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	u, err := s.h.GetUser(id)
	if err != nil {
		helper.DecodeError(w, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

// CreateUser ...
func (s *Setup) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqU := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(reqU); err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqU); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		helper.DecodeError(w, errPassHash, http.StatusInternalServerError)
		return
	}
	hashAPI := uuid.New()
	newU := model.User{
		Name:      reqU.Name,
		Email:     reqU.Email,
		Password:  hashPass,
		APIToken:  hashAPI,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	u, err := s.h.CreateUser(&newU)
	if err != nil {
		helper.DecodeError(w, errCreate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&u)
}

// UpdateUser ...
func (s *Setup) UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqU := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(reqU); err != nil {
		helper.DecodeError(w, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqU); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		helper.DecodeError(w, errPassHash, http.StatusInternalServerError)
		return
	}
	hashAPI := uuid.New()
	upU := model.User{
		Name:      reqU.Name,
		Email:     reqU.Email,
		Password:  hashPass,
		APIToken:  hashAPI,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	u, err := s.h.UpdateUser(id, &upU)
	if err != nil {
		helper.DecodeError(w, errUpdate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

// DeleteUser ...
func (s *Setup) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, errParseInt, http.StatusInternalServerError)
		return
	}
	if err := s.h.DeleteUser(id); err != nil {
		helper.DecodeError(w, errDelete, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "User deleted successfully!"})
}
