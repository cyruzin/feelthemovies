package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Genre type is a struct for genres table.
type Genre struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// GenreResult type is a result slice for genres.
type GenreResult struct {
	Data *[]Genre `json:"data"`
}

// GetGenres retrieves the latest genres.
func (c *Conn) GetGenres(ctx context.Context, limit int) (*GenreResult, error) {
	var result []Genre

	err := c.db.SelectContext(ctx, &result, queryGenresSelect, limit)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &GenreResult{&result}, nil
}

// GetGenre retrieves a genre by ID.
func (c *Conn) GetGenre(ctx context.Context, id int64) (*Genre, error) {
	var genre Genre

	err := c.db.GetContext(ctx, &genre, queryGenreSelectByID, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &genre, nil
}

// CreateGenre creates a new genre.
func (c *Conn) CreateGenre(ctx context.Context, g *Genre) error {
	_, err := c.db.ExecContext(ctx, queryGenreInsert, g.Name, g.CreatedAt, g.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGenre updates a genre by ID.
func (c *Conn) UpdateGenre(ctx context.Context, id int64, g *Genre) error {
	result, err := c.db.ExecContext(ctx, queryGenreUpdate, g.Name, g.UpdatedAt, id)
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

// DeleteGenre deletes a genre by ID.
func (c *Conn) DeleteGenre(ctx context.Context, id int64) error {
	result, err := c.db.ExecContext(ctx, queryGenreDelete, id)
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
