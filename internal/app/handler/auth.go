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
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	if err := s.validator.Struct(reqA); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	dbPass, err := s.model.Authenticate(reqA.Email)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errAuth, http.StatusInternalServerError)
		return
	}

	if checkPass := helper.CheckPasswordHash(reqA.Password, dbPass); !checkPass {
		helper.DecodeError(w, r, s.logger, errUnauthorized, http.StatusUnauthorized)
		return
	}

	authInfo, err := s.model.GetAuthInfo(reqA.Email)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	token, err := s.GenerateToken(authInfo)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	finalInfo := &model.AuthJWT{Token: token}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalInfo)
}

// GenerateToken generates a new JWT Token.
func (s *Setup) GenerateToken(info *model.Auth) (string, error) {
	secret := []byte(os.Getenv("JWTSECRET"))

	claims := struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		jwt.StandardClaims
	}{
		info.ID,
		info.Name,
		info.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Feel the Movies",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("Could not generate the Token")
	}
	return ss, nil
}
