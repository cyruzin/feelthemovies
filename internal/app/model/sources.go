package model

import (
	"database/sql"
	"errors"
	"time"
)

// Source type is a struct for sources table.
type Source struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// SourceResult type is a slice of sources.
type SourceResult struct {
	Data *[]Source `json:"data"`
}

// GetSources retrieves the latest 20 sources.
func (c *Conn) GetSources() (*SourceResult, error) {
	var sources []Source

	err := c.db.Select(&sources, querySourcesSelect, 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &SourceResult{&sources}, nil
}

// GetSource retrieves a source by a given ID.
func (c *Conn) GetSource(id int64) (*Source, error) {
	var source Source

	err := c.db.Get(&source, querySourcesSelectByID, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &source, nil
}

// CreateSource creates a new source.
func (c *Conn) CreateSource(s *Source) error {
	_, err := c.db.Exec(
		querySourcesInsert,
		s.Name,
		s.CreatedAt,
		s.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSource updates a source by a given ID.
func (c *Conn) UpdateSource(id int64, s *Source) error {
	result, err := c.db.Exec(
		querySourcesUpdate,
		s.Name,
		s.UpdatedAt,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}

// DeleteSource deletes a source by a given ID.
func (c *Conn) DeleteSource(id int64) error {
	result, err := c.db.Exec(querySourcesDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}
