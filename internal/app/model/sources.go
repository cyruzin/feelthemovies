package model

import (
	"database/sql"
	"log"
	"time"
)

// Source type is a struct for sources table.
type Source struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultSource type is a slice of sources.
type ResultSource []*Source

// GetSources retrieves the latest 20 sources.
func GetSources(db *sql.DB) (*ResultSource, error) {
	stmt, err := db.Query(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		ORDER BY id DESC
		LIMIT ?
	`, 20)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res := ResultSource{}

	for stmt.Next() {
		s := Source{}

		err = stmt.Scan(
			&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &s)

	}

	return &res, nil
}

// GetSource retrieves a source by a given ID.
func GetSource(id int64, db *sql.DB) (*Source, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		WHERE id = ?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	s := Source{}

	err = stmt.QueryRow(id).Scan(
		&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &s, nil
}

// CreateSource creates a new source.
func CreateSource(s *Source, db *sql.DB) (*Source, error) {
	stmt, err := db.Prepare(`
		INSERT INTO sources (
		name, created_at, updated_at
		)
		VALUES (?, ?, ?)
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&s.Name, &s.CreatedAt, &s.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetSource(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// UpdateSource updates a source by a given ID.
func UpdateSource(id int64, s *Source, db *sql.DB) (*Source, error) {
	stmt, err := db.Prepare(`
		UPDATE sources
		SET name=?, updated_at=?
		WHERE id=?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&s.Name, &s.UpdatedAt, &id,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetSource(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// DeleteSource deletes a source by a given ID.
func DeleteSource(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM sources
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
