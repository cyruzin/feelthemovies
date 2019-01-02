package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/gorilla/mux"
)

// TODO: Hash Password

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	u, err := model.GetUsers(db)

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

	u, err := model.GetUser(id, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqU)

	newU := model.User{
		Name:      reqU.Name,
		Email:     reqU.Email,
		Password:  reqU.Password,
		APIToken:  helper.RandStringRunes(32),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	u, err := model.CreateUser(&newU, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqU)

	upU := model.User{
		Name:      reqU.Name,
		Email:     reqU.Email,
		Password:  reqU.Password,
		APIToken:  reqU.APIToken,
		UpdatedAt: time.Now(),
	}

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	u, err := model.UpdateUser(id, &upU, db)

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

	d, err := model.DeleteUser(id, db)

	if d == 0 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("Something went wrong!")
	}

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Deleted Successfully!")
	}
}

func authUser(w http.ResponseWriter, r *http.Request) {

}
