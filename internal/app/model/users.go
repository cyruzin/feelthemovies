package model

import (
	"database/sql"
	"log"
	"time"
)

//TODO: Pagination

// User type is a struct for users table.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	APIToken  string    `json:"api_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultUser type is a slice of users.
type ResultUser []*User

// GetUsers retrieves the first twenty users.
func GetUsers(db *sql.DB) (*ResultUser, error) {

	stmt, err := db.Query(`
		SELECT 
		id, name, email, password,
		api_token, created_at, updated_at
		FROM users
		LIMIT 20
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res := ResultUser{}

	for stmt.Next() {
		user := User{}

		err = stmt.Scan(
			&user.ID, &user.Name, &user.Email, &user.Password,
			&user.APIToken, &user.CreatedAt, &user.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &user)

	}

	return &res, nil
}

// GetUser retrieves a user by a given ID.
func GetUser(id int64, db *sql.DB) (*User, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, email, password,
		api_token, created_at, updated_at
		FROM users
		WHERE id = ?
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	user := User{}

	err = stmt.QueryRow(id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password,
		&user.APIToken, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &user, nil
}

// CreateUser creates a new user.
func CreateUser(u *User, db *sql.DB) (*User, error) {
	stmt, err := db.Prepare(`
		INSERT INTO users (
		name, email, password,
		api_token, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&u.Name, &u.Email, &u.Password, &u.APIToken,
		&u.CreatedAt, &u.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetUser(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// UpdateUser updates a user by a given ID.
func UpdateUser(id int64, u *User, db *sql.DB) (*User, error) {
	stmt, err := db.Prepare(`
		UPDATE users
		SET name=?, email=?, password=?,
		api_token=?, updated_at=?
		WHERE id=?
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&u.Name, &u.Email, &u.Password, &u.APIToken,
		&u.UpdatedAt, &id,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetUser(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil

}

// DeleteUser deletes a user by a given ID.
func DeleteUser(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM users
		WHERE id=?
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	data, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
