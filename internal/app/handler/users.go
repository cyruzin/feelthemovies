package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// GetUsers get all users.
func (s *Setup) GetUsers(w http.ResponseWriter, r *http.Request) {
	u, err := s.h.GetUsers()
	if err != nil {
		helper.DecodeError(w,  s.l, err, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

// GetUser gets a user by ID.
func (s *Setup) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errParseInt, http.StatusInternalServerError)
		return
	}
	u, err := s.h.GetUser(id)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errFetch, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

// CreateUser creates a new user.
func (s *Setup) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqU := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(reqU); err != nil {
		helper.DecodeError(w,  s.l, err, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqU); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errPassHash, http.StatusInternalServerError)
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
		helper.DecodeError(w,  s.l, err, errCreate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&u)
}

// UpdateUser updates a user.
func (s *Setup) UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqU := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(reqU); err != nil {
		helper.DecodeError(w,  s.l, err, errDecode, http.StatusInternalServerError)
		return
	}
	if err := s.v.Struct(reqU); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errPassHash, http.StatusInternalServerError)
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
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errParseInt, http.StatusInternalServerError)
		return
	}
	u, err := s.h.UpdateUser(id, &upU)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errUpdate, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

// DeleteUser deletes a user.
func (s *Setup) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helper.DecodeError(w,  s.l, err, errParseInt, http.StatusInternalServerError)
		return
	}
	if err := s.h.DeleteUser(id); err != nil {
		helper.DecodeError(w,  s.l, err, errDelete, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "User deleted successfully!"})
}
