package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/pkg/helper"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getKeywords(w http.ResponseWriter, r *http.Request) {
	k, err := db.GetKeywords()
	if err != nil {
		helper.DecodeError(w, "Could not fetch the keywords", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&k)
}

func getKeyword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	k, err := db.GetKeyword(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the keyword", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&k)
}

func createKeyword(w http.ResponseWriter, r *http.Request) {
	reqK := &model.Keyword{}
	err := json.NewDecoder(r.Body).Decode(reqK)
	if err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqK); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	newK := model.Keyword{
		Name:      reqK.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	k, err := db.CreateKeyword(&newK)
	if err != nil {
		helper.DecodeError(w, "Could not create the keyword", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&k)
}

func updateKeyword(w http.ResponseWriter, r *http.Request) {
	reqK := &model.Keyword{}
	err := json.NewDecoder(r.Body).Decode(reqK)
	if err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqK); err != nil {
		helper.ValidatorMessage(w, err)
	}
	upK := model.Keyword{
		Name:      reqK.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	k, err := db.UpdateKeyword(id, &upK)
	if err != nil {
		helper.DecodeError(w, "Could not update the keyword", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&k)
}

func deleteKeyword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	if err := db.DeleteKeyword(id); err != nil {
		helper.DecodeError(w, "Could not delete the keyword", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "Keyword deleted successfully!"})
}
