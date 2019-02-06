package model

import (
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/go-sql-driver/mysql"
)

// User type is a struct for users table.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	APIToken  uuid.UUID `json:"api_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultUser type is a slice of users.
type ResultUser struct {
	Data []*User `json:"data"`
}

// GetUsers retrieves the first twenty users.
func (db *Conn) GetUsers() (*ResultUser, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, email, password,
		api_token, created_at, updated_at
		FROM users
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		res.Data = append(res.Data, &user)
	}
	return &res, nil
}

// GetUser retrieves a user by a given ID.
func (db *Conn) GetUser(id int64) (*User, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, email, password,
		api_token, created_at, updated_at
		FROM users
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	user := User{}
	err = stmt.QueryRow(id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password,
		&user.APIToken, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user.
func (db *Conn) CreateUser(u *User) (*User, error) {
	stmt, err := db.Prepare(`
		INSERT INTO users (
		name, email, password,
		api_token, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	data, err := db.GetUser(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateUser updates a user by a given ID.
func (db *Conn) UpdateUser(id int64, u *User) (*User, error) {
	stmt, err := db.Prepare(`
		UPDATE users
		SET name=?, email=?, password=?,
		api_token=?, updated_at=?
		WHERE id=?
	`)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	data, err := db.GetUser(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteUser deletes a user by a given ID.
func (db *Conn) DeleteUser(id int64) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM users
		WHERE id=?
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	data, err := res.RowsAffected()
	log.Println(data)
	if err != nil {
		return 0, err
	}
	return data, err
}