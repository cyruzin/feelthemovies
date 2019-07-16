package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Source type is a struct for sources table.
type Source struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SourceResult type is a slice of sources.
type SourceResult struct {
	Data []*Source `json:"data"`
}

// GetSources retrieves the latest 20 sources.
func (c *Conn) GetSources() (*SourceResult, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(10)
	if err != nil {
		return nil, err
	}
	res := SourceResult{}
	for rows.Next() {
		s := Source{}
		err = rows.Scan(
			&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		res.Data = append(res.Data, &s)
	}
	return &res, nil
}

// GetSource retrieves a source by a given ID.
func (c *Conn) GetSource(id int64) (*Source, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	s := Source{}
	err = stmt.QueryRow(id).Scan(
		&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &s, nil
}

// CreateSource creates a new source.
func (c *Conn) CreateSource(s *Source) (*Source, error) {
	stmt, err := c.db.Prepare(`
		INSERT INTO sources (
		name, created_at, updated_at
		)
		VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&s.Name, &s.CreatedAt, &s.UpdatedAt,
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
	data, err := c.GetSource(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateSource updates a source by a given ID.
func (c *Conn) UpdateSource(id int64, s *Source) (*Source, error) {
	stmt, err := c.db.Prepare(`
		UPDATE sources
		SET name=?, updated_at=?
		WHERE id=?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&s.Name, &s.UpdatedAt, &id,
	)
	if err != nil {
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	data, err := c.GetSource(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteSource deletes a source by a given ID.
func (c *Conn) DeleteSource(id int64) error {
	stmt, err := c.db.Prepare(`
		DELETE 
		FROM sources
		WHERE id=?
`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	data, err := res.RowsAffected()
	if err != nil || data == 0 {
		return errors.New("The resource you requested could not be found")
	}
	return nil
}
