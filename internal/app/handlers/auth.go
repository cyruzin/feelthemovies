package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

func authUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var reqU model.User

	err = json.NewDecoder(r.Body).Decode(&reqU)

	// TODO: Remove password field from this request.
	// This is only for comparison.
	user, err := model.Authenticate(reqU.Email, db)

	if err != nil {
		log.Println(err)
	}

	checkPass := helper.CheckPasswordHash(reqU.Password, user.Password)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else if !checkPass {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode("Unauthorized.")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)
	}
}
