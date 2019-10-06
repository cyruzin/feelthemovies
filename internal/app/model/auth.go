package model

import (
	"context"
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
func (c *Conn) Authenticate(ctx context.Context, email string) (string, error) {
	var password string

	err := c.db.GetContext(ctx, &password, queryAuthAuthenticate, email)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return password, nil
}

// GetAuthenticationInfo retrieves info for the authenticated user.
func (c *Conn) GetAuthenticationInfo(ctx context.Context, email string) (*Auth, error) {
	var auth Auth

	err := c.db.GetContext(ctx, &auth, queryAuthGetInfo, email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &auth, nil
}
