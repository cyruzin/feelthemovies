package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getSources(w http.ResponseWriter, r *http.Request) {
	s, err := db.GetSources()
	if err != nil {
		helper.DecodeError(w, "Could not fetch the sources", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

func getSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	s, err := db.GetSource(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the source", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

func createSource(w http.ResponseWriter, r *http.Request) {
	reqS := &model.Source{}
	if err := json.NewDecoder(r.Body).Decode(reqS); err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqS); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	newS := model.Source{
		Name:      reqS.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s, err := db.CreateSource(&newS)
	if err != nil {
		helper.DecodeError(w, "Could not create the source", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&s)
}

func updateSource(w http.ResponseWriter, r *http.Request) {
	reqS := &model.Source{}
	if err := json.NewDecoder(r.Body).Decode(reqS); err != nil {
		helper.DecodeError(w, "Could not decode the body response", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqS); err != nil {
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
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	s, err := db.UpdateSource(id, &upS)
	if err != nil {
		helper.DecodeError(w, "Could not update the source", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

func deleteSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	if err := db.DeleteSource(id); err != nil {
		helper.DecodeError(w, "Could not delete the source", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Source deleted successfully!"})
}
