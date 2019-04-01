package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	jwt "github.com/dgrijalva/jwt-go"
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

	token, err := s.GenerateToken()
	if err != nil {
		helper.DecodeError(w, r, s.l, errFetch, http.StatusInternalServerError)
		return
	}

	finalInfo := &model.AuthJWT{
		Auth:  authInfo,
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalInfo)
}

// GenerateToken generates a new JWT Token.
func (s *Setup) GenerateToken() (string, error) {
	secret := []byte(os.Getenv("JWTSecret"))

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    "Feel the Movies",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("Could not generate the Token")
	}
	return ss, nil
}
