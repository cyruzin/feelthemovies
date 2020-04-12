package models

import (
	"context"

	// Importing MySQL driver so the helper
	// functions can work.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	errResourceNotFound = "The resource you requested could not be found"
)

// Conn type is a struct for connections.
type Conn struct {
	db *sqlx.DB
}

// New connects to the database.
func New(db *sqlx.DB) *Conn {
	return &Conn{db}
}

// Attach receives a map of int/[]int and attach
// the IDs on the given pivot table.
func (c *Conn) Attach(ctx context.Context, s map[int64][]int, pivot string) error {
	for index, ids := range s {
		for _, values := range ids {
			query := "INSERT INTO " + pivot + " VALUES (?,?)"

			_, err := c.db.ExecContext(ctx, query, index, values)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Detach receives a map of int/[]int and Detach
// the IDs on the given pivot table.
func (c *Conn) Detach(ctx context.Context, s map[int64][]int, pivot, field string) error {
	for index := range s {
		query := "DELETE FROM " + pivot + " WHERE " + field + " = ?"

		_, err := c.db.ExecContext(ctx, query, index)
		if err != nil {
			return err
		}
	}

	return nil
}

// Sync receives a map of int/[]int and sync
// the IDs on the given pivot table.
func (c *Conn) Sync(ctx context.Context, s map[int64][]int, pivot, field string) error {
	empty := c.IsEmpty(s)

	if !empty {
		err := c.Detach(ctx, s, pivot, field)
		if err != nil {
			return err
		}

		err = c.Attach(ctx, s, pivot)
		if err != nil {
			return err
		}
	} else {
		err := c.Detach(ctx, s, pivot, field)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsEmpty checks if a given map of int/[]int is empty.
func (c *Conn) IsEmpty(s map[int64][]int) bool {
	empty := true

	for _, ids := range s {
		if len(ids) > 0 {
			empty = false
		}
	}

	return empty
}
