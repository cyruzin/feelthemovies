package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
)

// GetUsers get all users.
func (s *Setup) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.model.GetUsers(r.Context())
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &users)
}

// GetUser gets a user by ID.
func (s *Setup) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	user, err := s.model.GetUser(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &user)
}

// CreateUser creates a new user.
func (s *Setup) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	hashPass, err := s.HashPassword(request.Password, 10)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errPassHash, http.StatusInternalServerError)
		return
	}

	hashAPI := uuid.New()

	user := model.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashPass,
		APIToken:  hashAPI,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.model.CreateUser(ctx, &user)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &errhandler.APIMessage{Message: "User created successfully!"})
}

// UpdateUser updates a user.
func (s *Setup) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	hashPass, err := s.HashPassword(request.Password, 10)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errPassHash, http.StatusInternalServerError)
		return
	}

	hashAPI := uuid.New()

	user := model.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashPass,
		APIToken:  hashAPI,
		UpdatedAt: time.Now(),
	}

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.model.UpdateUser(ctx, id, &user)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "User updated successfully!"})
}

// DeleteUser deletes a user.
func (s *Setup) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteUser(r.Context(), id); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "User deleted successfully!"})
}
