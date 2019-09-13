package model

import (
	"database/sql"
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

// GetGenres retrieves the latest 10 genres.
func (c *Conn) GetGenres() (*GenreResult, error) {
	var result []Genre

	err := c.db.Select(&result, queryGenresSelect, 10)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &GenreResult{&result}, nil
}

// GetGenre retrieves a genre by ID.
func (c *Conn) GetGenre(id int64) (*Genre, error) {
	var genre Genre

	err := c.db.Get(&genre, queryGenreSelectByID, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &genre, nil
}

// CreateGenre creates a new genre.
func (c *Conn) CreateGenre(g *Genre) error {
	_, err := c.db.Exec(queryGenreInsert, g.Name, g.CreatedAt, g.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGenre updates a genre by ID.
func (c *Conn) UpdateGenre(id int64, g *Genre) error {
	_, err := c.db.Exec(queryGenreUpdate, g.Name, g.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGenre deletes a genre by ID.
func (c *Conn) DeleteGenre(id int64) error {
	_, err := c.db.Exec(queryGenreDelete, id)
	if err != nil {
		return err
	}

	return nil
}
