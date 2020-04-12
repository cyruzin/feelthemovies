package models

import (
	"context"
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
func (c *Conn) GetSources(ctx context.Context) (*SourceResult, error) {
	var sources []Source

	err := c.db.SelectContext(ctx, &sources, querySourcesSelect, 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &SourceResult{&sources}, nil
}

// GetSource retrieves a source by a given ID.
func (c *Conn) GetSource(ctx context.Context, id int64) (*Source, error) {
	var source Source

	err := c.db.GetContext(ctx, &source, querySourcesSelectByID, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &source, nil
}

// CreateSource creates a new source.
func (c *Conn) CreateSource(ctx context.Context, s *Source) error {
	_, err := c.db.ExecContext(
		ctx,
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
func (c *Conn) UpdateSource(ctx context.Context, id int64, s *Source) error {
	result, err := c.db.ExecContext(
		ctx,
		querySourcesUpdate,
		s.Name,
		s.UpdatedAt,
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

// DeleteSource deletes a source by a given ID.
func (c *Conn) DeleteSource(ctx context.Context, id int64) error {
	result, err := c.db.ExecContext(ctx, querySourcesDelete, id)
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
