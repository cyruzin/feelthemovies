package controllers

import (
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
	"github.com/go-chi/chi"

	model "github.com/cyruzin/feelthemovies/internal/app/models"
)

// GetSources gets all sources.
func (s *Setup) GetSources(w http.ResponseWriter, r *http.Request) {
	sources, err := s.model.GetSources(r.Context())
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &sources)
}

// GetSource gets a source by ID.
func (s *Setup) GetSource(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	source, err := s.model.GetSource(r.Context(), id)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &source)
}

// CreateSource creates a new source.
func (s *Setup) CreateSource(w http.ResponseWriter, r *http.Request) {
	request := model.Source{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	source := model.Source{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.model.CreateSource(ctx, &source)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &errhandler.APIMessage{Message: "Source created successfully!"})
}

// UpdateSource updates a source.
func (s *Setup) UpdateSource(w http.ResponseWriter, r *http.Request) {
	request := model.Source{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	source := model.Source{
		Name:      request.Name,
		UpdatedAt: time.Now(),
	}

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.model.UpdateSource(ctx, id, &source)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "Source updated successfully!"})
}

// DeleteSource deletes a source.
func (s *Setup) DeleteSource(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteSource(r.Context(), id); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &errhandler.APIMessage{Message: "Source deleted successfully!"})
}
