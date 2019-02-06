package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/cyruzin/feelthemovies/pkg/helper"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	u, err := db.GetUsers()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(u)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	u, err := db.GetUser(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(u)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqU model.User
	err := json.NewDecoder(r.Body).Decode(&reqU)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqU)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		log.Println(err)
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
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(u)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqU model.User
	err := json.NewDecoder(r.Body).Decode(&reqU)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqU)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	hashPass, err := helper.HashPassword(reqU.Password, 10)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	u, err := db.UpdateUser(id, &upU)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(u)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	d, err := db.DeleteUser(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else if d == 0 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Deleted Successfully!")
	}
}