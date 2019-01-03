package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Keyword type is a struct for keywords table.
type Keyword struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultKeyword type is a slice of keywords.
type ResultKeyword struct {
	Data []*Keyword `json:"data"`
}

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
		log.Println(err)
	}

	defer stmt.Close()

	res := ResultKeyword{}

	for stmt.Next() {
		k := Keyword{}

		err = stmt.Scan(
			&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &k)

	}

	return &res, err
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
		log.Println(err)
	}

	defer stmt.Close()

	k := Keyword{}

	err = stmt.QueryRow(id).Scan(
		&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	return &k, err
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
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&k.Name, &k.CreatedAt, &k.UpdatedAt,
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

	data, err := GetKeyword(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// UpdateKeyword updates a keyword by a given ID.
func UpdateKeyword(id int64, k *Keyword, db *sql.DB) (*Keyword, error) {
	stmt, err := db.Prepare(`
		UPDATE keywords
		SET name=?, updated_at=?
		WHERE id=?
`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&k.Name, &k.UpdatedAt, &id,
	)

	if err != nil {
		log.Println(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	data, err := GetKeyword(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// DeleteKeyword deletes a keyword by a given ID.
func DeleteKeyword(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM keywords
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
