package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getSources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	s, err := db.GetSources()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(s)
	}
}

func getSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	s, err := db.GetSource(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(s)
	}
}

func createSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqS model.Source
	err := json.NewDecoder(r.Body).Decode(&reqS)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqS)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	newS := model.Source{
		Name:      reqS.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s, err := db.CreateSource(&newS)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(s)
	}
}

func updateSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqS model.Source
	err := json.NewDecoder(r.Body).Decode(&reqS)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqS)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	upS := model.Source{
		Name:      reqS.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	s, err := db.UpdateSource(id, &upS)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(s)
	}
}

func deleteSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	d, err := db.DeleteSource(id)
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
