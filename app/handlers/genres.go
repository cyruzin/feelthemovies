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

func getGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	g, err := db.GetGenres()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(g)
	}
}

func getGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	g, err := db.GetGenre(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(g)
	}

}

func createGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqG model.Genre
	err := json.NewDecoder(r.Body).Decode(&reqG)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqG)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	newG := model.Genre{
		Name:      reqG.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	g, err := db.CreateGenre(&newG)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(g)
	}
}

func updateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqG model.Genre
	err := json.NewDecoder(r.Body).Decode(&reqG)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqG)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	upG := model.Genre{
		Name:      reqG.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	g, err := db.UpdateGenre(id, &upG)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(g)
	}
}

func deleteGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	d, err := db.DeleteGenre(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else if d == 0 {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("The resource you requested could not be found.")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Deleted Successfully!")
	}
}
