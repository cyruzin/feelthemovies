package model

import (
	"database/sql"
	"log"
	"time"
)

// Keyword type is a struct for keywords table.
type Keyword struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultKeyword type is a slice of keywords.
type ResultKeyword []*Keyword

// GetKeywords retrieves the latest 20 keywords.
func GetKeywords(db *sql.DB) (*ResultKeyword, error) {
	stmt, err := db.Query(`
		SELECT 
		id, name, created_at, updated_at
		FROM keywords
		ORDER BY id DESC
		LIMIT ?
	`, 20)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res := ResultKeyword{}

	for stmt.Next() {
		k := Keyword{}

		err = stmt.Scan(
			&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &k)

	}

	return &res, nil
}

// GetKeyword retrieves a keyword by a given ID.
func GetKeyword(id int64, db *sql.DB) (*Keyword, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM keywords
		WHERE id = ?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	k := Keyword{}

	err = stmt.QueryRow(id).Scan(
		&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &k, nil
}

// CreateKeyword creates a new keyword.
func CreateKeyword(k *Keyword, db *sql.DB) (*Keyword, error) {
	stmt, err := db.Prepare(`
		INSERT INTO keywords (
		name, created_at, updated_at
		)
		VALUES (?, ?, ?)
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&k.Name, &k.CreatedAt, &k.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetKeyword(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// UpdateKeyword updates a keyword by a given ID.
func UpdateKeyword(id int64, k *Keyword, db *sql.DB) (*Keyword, error) {
	stmt, err := db.Prepare(`
		UPDATE keywords
		SET name=?, updated_at=?
		WHERE id=?
`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&k.Name, &k.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetKeyword(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// DeleteKeyword deletes a keyword by a given ID.
func DeleteKeyword(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM keywords
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
