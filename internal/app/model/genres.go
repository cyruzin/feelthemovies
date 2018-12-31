package model

import (
	"database/sql"
	"log"
	"time"
)

// Genre type is a struct for genres table.
type Genre struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultGenre type is a slice of genres.
type ResultGenre []*Genre

// GetGenres retrieves the latest 20 genres.
func GetGenres(db *sql.DB) (*ResultGenre, error) {
	stmt, err := db.Query(`
		SELECT 
		id, name, created_at, updated_at
		FROM genres
		ORDER BY id DESC
		LIMIT ?
	`, 20)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res := ResultGenre{}

	for stmt.Next() {
		genre := Genre{}

		err = stmt.Scan(
			&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &genre)

	}

	return &res, nil
}

// GetGenre retrieves a genre by a given ID.
func GetGenre(id int64, db *sql.DB) (*Genre, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM genres
		WHERE id = ?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	genre := Genre{}

	err = stmt.QueryRow(id).Scan(
		&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &genre, nil
}

// CreateGenre creates a new genre.
func CreateGenre(g *Genre, db *sql.DB) (*Genre, error) {
	stmt, err := db.Prepare(`
		INSERT INTO genres (
		name, created_at, updated_at
		)
		VALUES (?, ?, ?)
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&g.Name, &g.CreatedAt, &g.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetGenre(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// UpdateGenre updates a genre by a given ID.
func UpdateGenre(id int64, g *Genre, db *sql.DB) (*Genre, error) {
	stmt, err := db.Prepare(`
		UPDATE genres
		SET name=?, updated_at=?
		WHERE id=?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&g.Name, &g.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetGenre(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// DeleteGenre deletes a genre by a given ID.
func DeleteGenre(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM genres
		WHERE id=?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	data, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
