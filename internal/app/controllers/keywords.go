package controllers

import (
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
	"github.com/go-chi/chi"

	model "github.com/cyruzin/feelthemovies/internal/app/models"
)

// GetKeywords gets all keywords.
func (s *Setup) GetKeywords(w http.ResponseWriter, r *http.Request) {
	keywords, err := s.model.GetKeywords(r.Context(), 20)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &keywords)
}

// GetKeyword gets a keyword by ID.
func (s *Setup) GetKeyword(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	keyword, err := s.model.GetKeyword(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &keyword)
}

// CreateKeyword creates a new keyword.
func (s *Setup) CreateKeyword(w http.ResponseWriter, r *http.Request) {
	keyword := model.Keyword{}

	err := json.NewDecoder(r.Body).Decode(&keyword)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, keyword); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	keyword.CreatedAt = time.Now()
	keyword.UpdatedAt = time.Now()

	err = s.model.CreateKeyword(ctx, &keyword)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &errhandler.APIMessage{Message: "Keyword created successfully!"})
}

// UpdateKeyword updates a keyword.
func (s *Setup) UpdateKeyword(w http.ResponseWriter, r *http.Request) {
	keyword := model.Keyword{}

	err := json.NewDecoder(r.Body).Decode(&keyword)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, keyword); err != nil {
		validation.ValidatorMessage(w, err)
	}

	keyword.UpdatedAt = time.Now()

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.model.UpdateKeyword(ctx, id, &keyword)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "Keyword updated successfully!"})
}

// DeleteKeyword deletes a keyword.
func (s *Setup) DeleteKeyword(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteKeyword(r.Context(), id); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "Keyword deleted successfully!"})
}
