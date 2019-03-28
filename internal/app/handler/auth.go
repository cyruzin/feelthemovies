package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

// AuthUser authenticates the user.
func (s *Setup) AuthUser(w http.ResponseWriter, r *http.Request) {
	var reqA model.Auth

	if err := json.NewDecoder(r.Body).Decode(&reqA); err != nil {
		helper.DecodeError(w, r, s.l, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.v.Struct(reqA); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	dbPass, err := s.h.Authenticate(reqA.Email)
	if err != nil {
		helper.DecodeError(w, r, s.l, errAuth, http.StatusInternalServerError)
		return
	}

	if checkPass := helper.CheckPasswordHash(reqA.Password, dbPass); !checkPass {
		helper.DecodeError(w, r, s.l, errUnauthorized, http.StatusUnauthorized)
		return
	}

	authInfo, err := s.h.GetAuthInfo(reqA.Email)
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&authInfo)
}
