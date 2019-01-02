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

// Authenticate the current user and returns it's info.
// TODO: Implement! Use User struct.
func Authenticate() bool {
	return true
}
