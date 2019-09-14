package model

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	errResourceNotFound = "The resource you requested could not be found"
)

// Conn type is a struct for connections.
type Conn struct {
	db *sqlx.DB
}

// Connect connects to the database.
func Connect(db *sqlx.DB) *Conn {
	return &Conn{db}
}

// Attach receives a map of int/[]int and attach
// the IDs on the given pivot table.
func (c *Conn) Attach(s map[int64][]int, pivot string) error {
	for index, ids := range s {
		for _, values := range ids {
			query := "INSERT INTO " + pivot + " VALUES (?,?)"
			stmt, err := c.db.Prepare(query)
			if err != nil {
				return err
			}
			defer stmt.Close()
			_, err = stmt.Exec(index, values)
			// Error handler for duplicate entries
			if mysqlError, ok := err.(*mysql.MySQLError); ok {
				if mysqlError.Number == 1062 {
					return err
				}
			}
		}
	}
	return nil
}

// Detach receives a map of int/[]int and Detach
// the IDs on the given pivot table.
func (c *Conn) Detach(s map[int64][]int, pivot, field string) error {
	for index := range s {
		query := "DELETE FROM " + pivot + " WHERE " + field + " = ?"
		stmt, err := c.db.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(index)
		if err != nil {
			return err
		}
	}
	return nil
}

// Sync receives a map of int/[]int and sync
// the IDs on the given pivot table.
func (c *Conn) Sync(s map[int64][]int, pivot, field string) error {
	empty := c.IsEmpty(s)
	if !empty {
		err := c.Detach(s, pivot, field)
		if err != nil {
			return err
		}
		err = c.Attach(s, pivot)
		if err != nil {
			return err
		}
	} else {
		err := c.Detach(s, pivot, field)
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
