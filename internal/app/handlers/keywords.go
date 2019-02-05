package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

func getKeywords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	k, err := db.GetKeywords()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(k)
	}
}

func getKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	k, err := db.GetKeyword(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(k)
	}

}

func createKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqK model.Keyword
	err := json.NewDecoder(r.Body).Decode(&reqK)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqK)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	newK := model.Keyword{
		Name:      reqK.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	k, err := db.CreateKeyword(&newK)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(k)
	}
}

func updateKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqK model.Keyword
	err := json.NewDecoder(r.Body).Decode(&reqK)
	if err != nil {
		log.Println(err)
	}
	validate = validator.New()
	err = validate.Struct(reqK)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	upK := model.Keyword{
		Name:      reqK.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	k, err := db.UpdateKeyword(id, &upK)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(k)
	}
}

func deleteKeyword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	d, err := db.DeleteKeyword(id)
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
