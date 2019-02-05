package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"
)

func authUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var reqA model.Auth
	err = json.NewDecoder(r.Body).Decode(&reqA)
	validate = validator.New()
	err = validate.Struct(reqA)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Validation error, check your fields.")
		return
	}
	dbPass, err := db.Authenticate(reqA.Email)
	if err != nil {
		log.Println(err)
	}
	checkPass := helper.CheckPasswordHash(reqA.Password, dbPass)
	if !checkPass {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode("Unauthorized.")
		return
	}
	authInfo, err := db.GetAuthInfo(reqA.Email)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(authInfo)
	}
}
