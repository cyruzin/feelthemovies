package handler

import (
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
	request := model.Auth{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		helper.ValidatorMessage(w, err)
		return
	}

	dbPassword, err := s.model.Authenticate(ctx, request.Email)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errAuth, http.StatusInternalServerError)
		return
	}

	if checkPassword := helper.CheckPasswordHash(request.Password, dbPassword); !checkPassword {
		helper.DecodeError(w, r, s.logger, errUnauthorized, http.StatusUnauthorized)
		return
	}

	authenticationInfo, err := s.model.GetAuthenticationInfo(ctx, request.Email)
	if err != nil {
		helper.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	token, err := s.GenerateToken(authenticationInfo)
	if err != nil {
		helper.DecodeError(w, r, s.logger, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo := model.AuthJWT{Token: token}

	s.ToJSON(w, http.StatusOK, &userInfo)
}

// GenerateToken generates a new JWT Token.
func (s *Setup) GenerateToken(info *model.Auth) (string, error) {
	secret := []byte(os.Getenv("JWTSECRET"))

	var claims model.AuthClaims

	claims.ID = info.ID
	claims.Name = info.Name
	claims.Email = info.Email
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Issuer:    "Feel the Movies",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New(errGeneratingToken)
	}

	return signedString, nil
}
