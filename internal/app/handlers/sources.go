package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

func getSources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	s, err := model.GetSources(db)

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

	s, err := model.GetSource(id, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqS)

	newS := model.Source{
		Name:      reqS.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s, err := model.CreateSource(&newS, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqS)

	upS := model.Source{
		Name:      reqS.Name,
		UpdatedAt: time.Now(),
	}

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	s, err := model.UpdateSource(id, &upS, db)

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

	d, err := model.DeleteSource(id, db)

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
