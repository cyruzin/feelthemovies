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

func getGenres(w http.ResponseWriter, r *http.Request) {
	g, err := db.GetGenres()
	if err != nil {
		helper.DecodeError(w, "Could not fetch the genres", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&g)
}

func getGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	g, err := db.GetGenre(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the genre", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&g)
}

func createGenre(w http.ResponseWriter, r *http.Request) {
	reqG := &model.Genre{}
	err := json.NewDecoder(r.Body).Decode(reqG)
	if err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqG); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	newG := model.Genre{
		Name:      reqG.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	g, err := db.CreateGenre(&newG)
	if err != nil {
		helper.DecodeError(w, "Could not create the genre", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&g)
}

func updateGenre(w http.ResponseWriter, r *http.Request) {
	reqG := &model.Genre{}
	err := json.NewDecoder(r.Body).Decode(reqG)
	if err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqG); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	upG := model.Genre{
		Name:      reqG.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	g, err := db.UpdateGenre(id, &upG)
	if err != nil {
		helper.DecodeError(w, "Could not update the genre", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&g)
}

func deleteGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	if err := db.DeleteGenre(id); err != nil {
		helper.DecodeError(w, "Could not delete the genre", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Genre deleted successfully!"})
}
