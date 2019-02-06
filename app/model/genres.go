package model

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// Genre type is a struct for genres table.
type Genre struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultGenre type is a slice of genres.
type ResultGenre struct {
	Data []*Genre `json:"data"`
}

// GetGenres retrieves the latest 20 genres.
func (db *Conn) GetGenres() (*ResultGenre, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM genres
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(10)
	res := ResultGenre{}
	for rows.Next() {
		genre := Genre{}
		err = rows.Scan(
			&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &genre)
	}
	return &res, nil
}

// GetGenre retrieves a genre by a given ID.
func (db *Conn) GetGenre(id int64) (*Genre, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM genres
		WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	genre := Genre{}
	err = stmt.QueryRow(id).Scan(
		&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &genre, nil
}

// CreateGenre creates a new genre.
func (db *Conn) CreateGenre(g *Genre) (*Genre, error) {
	stmt, err := db.Prepare(`
		INSERT INTO genres (
		name, created_at, updated_at
		)
		VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&g.Name, &g.CreatedAt, &g.UpdatedAt,
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
	data, err := db.GetGenre(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateGenre updates a genre by a given ID.
func (db *Conn) UpdateGenre(id int64, g *Genre) (*Genre, error) {
	stmt, err := db.Prepare(`
		UPDATE genres
		SET name=?, updated_at=?
		WHERE id=?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&g.Name, &g.UpdatedAt, &id,
	)
	if err != nil {
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	data, err := db.GetGenre(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteGenre deletes a genre by a given ID.
func (db *Conn) DeleteGenre(id int64) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM genres
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
	if err != nil {
		return 0, err
	}
	return data, nil
}
