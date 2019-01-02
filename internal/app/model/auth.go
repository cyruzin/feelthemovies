package model

import (
	"database/sql"
	"log"
)

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
func Authenticate(email string, db *sql.DB) (*User, error) {
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

	var u User

	err = stmt.QueryRow(email).Scan(
		&u.ID, &u.Name, &u.Email, &u.Password, &u.APIToken,
	)

	if err != nil {
		log.Println(err)
	}

	return &u, err
}
