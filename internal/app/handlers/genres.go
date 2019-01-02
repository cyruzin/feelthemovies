package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

func getGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	g, err := model.GetGenres(db)

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

	g, err := model.GetGenre(id, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqG)

	newG := model.Genre{
		Name:      reqG.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	g, err := model.CreateGenre(&newG, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqG)

	upG := model.Genre{
		Name:      reqG.Name,
		UpdatedAt: time.Now(),
	}
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	g, err := model.UpdateGenre(id, &upG, db)

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

	d, err := model.DeleteGenre(id, db)

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
