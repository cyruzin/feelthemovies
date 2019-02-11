package model

import (
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Keyword type is a struct for keywords table.
type Keyword struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultKeyword type is a slice of keywords.
type ResultKeyword struct {
	Data []*Keyword `json:"data"`
}

// GetKeywords retrieves the latest 20 keywords.
func (db *Conn) GetKeywords() (*ResultKeyword, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM keywords
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(10)
	res := ResultKeyword{}
	for rows.Next() {
		k := Keyword{}
		err = rows.Scan(
			&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &k)
	}
	return &res, nil
}

// GetKeyword retrieves a keyword by a given ID.
func (db *Conn) GetKeyword(id int64) (*Keyword, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, name, created_at, updated_at
		FROM keywords
		WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	k := Keyword{}
	err = stmt.QueryRow(id).Scan(
		&k.ID, &k.Name, &k.CreatedAt, &k.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// CreateKeyword creates a new keyword.
func (db *Conn) CreateKeyword(k *Keyword) (*Keyword, error) {
	stmt, err := db.Prepare(`
		INSERT INTO keywords (
		name, created_at, updated_at
		)
		VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	data, err := db.GetKeyword(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateKeyword updates a keyword by a given ID.
func (db *Conn) UpdateKeyword(id int64, k *Keyword) (*Keyword, error) {
	stmt, err := db.Prepare(`
		UPDATE keywords
		SET name=?, updated_at=?
		WHERE id=?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&k.Name, &k.UpdatedAt, &id,
	)
	if err != nil {
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	data, err := db.GetKeyword(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteKeyword deletes a keyword by a given ID.
func (db *Conn) DeleteKeyword(id int64) error {
	stmt, err := db.Prepare(`
		DELETE 
		FROM keywords
		WHERE id=?
`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	data, err := res.RowsAffected()
	if err != nil || data == 0 {
		return errors.New("The resource you requested could not be found")
	}
	return nil
}
