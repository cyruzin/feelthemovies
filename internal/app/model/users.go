package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// User type is a struct for users table.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	APIToken  string    `json:"api_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultUser type is a slice of users.
type ResultUser struct {
	Data []*User `json:"data"`
}

// GetUsers retrieves the first twenty users.
func GetUsers(db *sql.DB) (*ResultUser, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, name, email, password,
		api_token, created_at, updated_at
		FROM users
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(10)

	res := ResultUser{}

	for rows.Next() {
		user := User{}

		err = rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Password,
			&user.APIToken, &user.CreatedAt, &user.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &user)

	}

	return &res, err
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
		log.Println(err)
	}

	defer stmt.Close()

	user := User{}

	err = stmt.QueryRow(id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password,
		&user.APIToken, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	return &user, err
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
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&u.Name, &u.Email, &u.Password, &u.APIToken,
		&u.CreatedAt, &u.UpdatedAt,
	)

	// Error handler for duplicate entries
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return nil, err
		}
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	data, err := GetUser(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
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
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&u.Name, &u.Email, &u.Password, &u.APIToken,
		&u.UpdatedAt, &id,
	)

	// Error handler for duplicate entries
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return nil, err
		}
	}

	if err != nil {
		log.Println(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	data, err := GetUser(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err

}

// DeleteUser deletes a user by a given ID.
func DeleteUser(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM users
		WHERE id=?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		log.Println(err)
	}

	data, err := res.RowsAffected()

	log.Println(data)

	if err != nil {
		log.Println(err)
	}

	return data, err
}
