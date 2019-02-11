package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/feelthemovies/app/model"
	"github.com/cyruzin/feelthemovies/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"
)

func authUser(w http.ResponseWriter, r *http.Request) {
	var reqA model.Auth

	if err := json.NewDecoder(r.Body).Decode(&reqA); err != nil {
		helper.DecodeError(w, "Could not decode the auth payload", http.StatusInternalServerError)
		return
	}

	validate = validator.New()
	if err := validate.Struct(reqA); err != nil {
		helper.DecodeError(w, "Validation error, check your fields", http.StatusBadRequest)
		return
	}

	dbPass, err := db.Authenticate(reqA.Email)
	if err != nil {
		helper.DecodeError(w, "Could not authenticate", http.StatusInternalServerError)
		return
	}

	if checkPass := helper.CheckPasswordHash(reqA.Password, dbPass); !checkPass {
		helper.DecodeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	authInfo, err := db.GetAuthInfo(reqA.Email)
	if err != nil {
		helper.DecodeError(w, "Could not get user info", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&authInfo)
}
