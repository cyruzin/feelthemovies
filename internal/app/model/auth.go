package model

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
)

// Auth type is a struct for authentication.
type Auth struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

// AuthClaims type is a struct for JWT Claims.
type AuthClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// AuthJWT type is a struct for JWT authentication.
type AuthJWT struct {
	Token string `json:"token"`
}

// Authenticate authenticates the current user and returns it's info.
func (c *Conn) Authenticate(email string) (string, error) {
	var password string

	err := c.db.Get(&password, queryAuthAuthenticate, email)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return password, nil
}

// GetAuthenticationInfo retrieves info for the authenticated user.
func (c *Conn) GetAuthenticationInfo(email string) (*Auth, error) {
	var auth Auth

	err := c.db.Get(&auth, queryAuthGetInfo, email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &auth, nil
}
