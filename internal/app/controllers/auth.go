package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/config"
	model "github.com/cyruzin/feelthemovies/internal/app/models"
	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/validation"
	jwt "github.com/dgrijalva/jwt-go"
)

// AuthUser authenticates the user.
func (s *Setup) AuthUser(w http.ResponseWriter, r *http.Request) {
	request := model.Auth{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errhandler.DecodeError(w, r, s.logger, errDecode, http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	if err := s.validator.StructCtx(ctx, request); err != nil {
		validation.ValidatorMessage(w, err)
		return
	}

	dbPassword, err := s.model.Authenticate(ctx, request.Email)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errAuth, http.StatusInternalServerError)
		return
	}

	if checkPassword := s.CheckPasswordHash(request.Password, dbPassword); !checkPassword {
		errhandler.DecodeError(w, r, s.logger, errUnauthorized, http.StatusUnauthorized)
		return
	}

	authenticationInfo, err := s.model.GetAuthenticationInfo(ctx, request.Email)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, errFetch, http.StatusInternalServerError)
		return
	}

	token, err := s.GenerateToken(authenticationInfo)
	if err != nil {
		errhandler.DecodeError(w, r, s.logger, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo := model.AuthJWT{Token: token}

	s.ToJSON(w, http.StatusOK, &userInfo)
}

// GenerateToken generates a new JWT Token.
func (s *Setup) GenerateToken(info *model.Auth) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}

	secret := []byte(cfg.JWTSecret)

	var claims model.AuthClaims

	claims.ID = info.ID
	claims.Name = info.Name
	claims.Email = info.Email
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Issuer:    cfg.AppName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New(errGeneratingToken)
	}

	return signedString, nil
}
