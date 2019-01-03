package model

import (
	"database/sql"
	"log"
)

// Auth struct is a type for authentication.
type Auth struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	APIToken string `json:"api_token"`
}

// CheckAPIToken checks if the given token exists.
func CheckAPIToken(token string, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(`
		SELECT 
		api_token
		FROM users
		WHERE api_token = ?
`)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	var t string

	err = stmt.QueryRow(token).Scan(&t)

	if err != nil {
		log.Println(err)
	}

	if t != "" && t == token {
		return true, err
	}

	return false, err
}

// Authenticate authenticates the current user and returns it's info.
func Authenticate(email string, db *sql.DB) (*Auth, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, email, password, api_token
		FROM users
		WHERE email = ?
`)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	var a Auth

	err = stmt.QueryRow(email).Scan(
		&a.ID, &a.Name, &a.Email, &a.Password, &a.APIToken,
	)

	if err != nil {
		log.Println(err)
	}

	return &a, err
}
