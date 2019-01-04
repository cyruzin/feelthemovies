package model

import (
	"database/sql"
	"log"
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

// ResultSource type is a slice of sources.
type ResultSource struct {
	Data []*Source `json:"data"`
}

// GetSources retrieves the latest 20 sources.
func GetSources(db *sql.DB) (*ResultSource, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM sources
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(10)

	res := ResultSource{}

	for rows.Next() {
		s := Source{}

		err = rows.Scan(
			&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &s)

	}

	return &res, err
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
		log.Println(err)
	}

	defer stmt.Close()

	s := Source{}

	err = stmt.QueryRow(id).Scan(
		&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	return &s, err
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
		log.Println(err)
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
		log.Println(err)
	}

	data, err := GetSource(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// UpdateSource updates a source by a given ID.
func UpdateSource(id int64, s *Source, db *sql.DB) (*Source, error) {
	stmt, err := db.Prepare(`
		UPDATE sources
		SET name=?, updated_at=?
		WHERE id=?
`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&s.Name, &s.UpdatedAt, &id,
	)

	if err != nil {
		log.Println(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	data, err := GetSource(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// DeleteSource deletes a source by a given ID.
func DeleteSource(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM sources
		WHERE id=?
`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		log.Println(err)
	}

	data, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	return data, err
}
