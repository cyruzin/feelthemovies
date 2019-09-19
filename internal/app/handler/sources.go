package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/go-chi/chi"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

// GetSources gets all sources.
func (s *Setup) GetSources(w http.ResponseWriter, r *http.Request) {
	sources, err := s.model.GetSources()
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &sources)
}

// GetSource gets a source by ID.
func (s *Setup) GetSource(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	source, err := s.model.GetSource(id)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &source)
}

// CreateSource creates a new source.
func (s *Setup) CreateSource(w http.ResponseWriter, r *http.Request) {
	request := model.Source{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	source := model.Source{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.model.CreateSource(&source)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errCreate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusCreated, &helper.APIMessage{Message: "Source created successfully!"})
}

// UpdateSource updates a source.
func (s *Setup) UpdateSource(w http.ResponseWriter, r *http.Request) {
	request := model.Source{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	source := model.Source{
		Name:      request.Name,
		UpdatedAt: time.Now(),
	}

	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	err = s.model.UpdateSource(id, &source)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errUpdate, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &helper.APIMessage{Message: "Source updated successfully!"})
}

// DeleteSource deletes a source.
func (s *Setup) DeleteSource(w http.ResponseWriter, r *http.Request) {
	id, err := s.IDParser(chi.URLParam(r, "id"))
	if err != nil {
		helper.DecodeError(w, r, s.logger, errParseInt, http.StatusInternalServerError)
		return
	}

	if err := s.model.DeleteSource(id); err != nil {
		helper.DecodeError(w, r, s.logger, errDelete, http.StatusInternalServerError)
		return
	}

	s.ToJSON(w, http.StatusOK, &helper.APIMessage{Message: "Source deleted successfully!"})
}
