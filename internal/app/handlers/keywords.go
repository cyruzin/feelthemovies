package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/gorilla/mux"
)

func getKeywords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	k, err := model.GetKeywords(db)

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

	k, err := model.GetKeyword(id, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqK)

	newK := model.Keyword{
		Name:      reqK.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	k, err := model.CreateKeyword(&newK, db)

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

	err = json.NewDecoder(r.Body).Decode(&reqK)

	upK := model.Keyword{
		Name:      reqK.Name,
		UpdatedAt: time.Now(),
	}

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	k, err := model.UpdateKeyword(id, &upK, db)

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

	d, err := model.DeleteKeyword(id, db)

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
