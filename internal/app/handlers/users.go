package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	u, err := db.GetUsers()
	if err != nil {
		helper.DecodeError(w, "Could not fetch the users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	u, err := db.GetUser(id)
	if err != nil {
		helper.DecodeError(w, "Could not fetch the user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	reqU := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(reqU); err != nil {
		helper.DecodeError(w, "Could not decode the body request", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqU); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		helper.DecodeError(w, "Could not hash the password", http.StatusInternalServerError)
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
	u, err := db.CreateUser(&newU)
	if err != nil {
		helper.DecodeError(w, "Could not create the user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	reqU := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(reqU); err != nil {
		helper.DecodeError(w, "Could not decode the body response", http.StatusInternalServerError)
		return
	}
	validate = validator.New()
	if err := validate.Struct(reqU); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		helper.DecodeError(w, "Could not hash the password", http.StatusInternalServerError)
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
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	u, err := db.UpdateUser(id, &upU)
	if err != nil {
		helper.DecodeError(w, "Could not update the user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		helper.DecodeError(w, "Could not parse the ID param", http.StatusInternalServerError)
		return
	}
	if err := db.DeleteUser(id); err != nil {
		helper.DecodeError(w, "Could not delete the user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&helper.APIMessage{Message: "User deleted successfully!"})
}
