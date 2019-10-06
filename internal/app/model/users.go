package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// User type is a struct for users table.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	APIToken  uuid.UUID `db:"api_token" json:"api_token"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// UserResult type is a slice of users.
type UserResult struct {
	Data *[]User `json:"data"`
}

// GetUsers retrieves the first twenty users.
func (c *Conn) GetUsers(ctx context.Context) (*UserResult, error) {
	var users []User

	err := c.db.SelectContext(ctx, &users, queryUsersSelect, 10)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &UserResult{&users}, nil
}

// GetUser retrieves a user by a given ID.
func (c *Conn) GetUser(ctx context.Context, id int64) (*User, error) {
	var user User

	err := c.db.GetContext(ctx, &user, queryUserSelectByID, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user.
func (c *Conn) CreateUser(ctx context.Context, u *User) error {
	_, err := c.db.ExecContext(
		ctx,
		queryUserInsert,
		u.Name,
		u.Email,
		u.Password,
		u.APIToken,
		u.CreatedAt,
		u.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user by a given ID.
func (c *Conn) UpdateUser(ctx context.Context, id int64, u *User) error {
	result, err := c.db.ExecContext(
		ctx,
		queryUserUpdate,
		u.Name,
		u.Email,
		u.Password,
		u.APIToken,
		u.UpdatedAt,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}

// DeleteUser deletes a user by a given ID.
func (c *Conn) DeleteUser(ctx context.Context, id int64) error {
	result, err := c.db.ExecContext(ctx, queryUserDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}
