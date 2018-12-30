package model

import (
	"database/sql"
	"log"
	"time"
)

// Recommendation type is a struct for recommendations.
type Recommendation struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Type      int       `json:"type"`
	Body      string    `json:"body"`
	Poster    string    `json:"poster"`
	Backdrop  string    `json:"backdrop"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Result type is a slice of recommendation.
type Result []*Recommendation

// GetRecommendations retrieves the latest 20 recommendations.
func GetRecommendations(db *sql.DB) (*Result, error) {

	stmtOut, err := db.Query(`
		SELECT 
		id, user_id, title, type, 
		body, poster, backdrop, status, 
		created_at, updated_at
		FROM recommendations
		ORDER BY id DESC
		LIMIT 20
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmtOut.Close()

	res := Result{}

	for stmtOut.Next() {
		rec := Recommendation{}

		err = stmtOut.Scan(
			&rec.ID, &rec.UserID, &rec.Title, &rec.Type,
			&rec.Body, &rec.Backdrop, &rec.Poster, &rec.Status,
			&rec.CreatedAt, &rec.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &rec)

	}

	return &res, nil
}

// GetRecommendation retrieves a recommendation by a given ID.
func GetRecommendation(id int64, db *sql.DB) (*Recommendation, error) {
	stmtOut, err := db.Prepare(`
		SELECT 
		id, user_id, title, type, body, poster, 
		backdrop, status, created_at, updated_at
		FROM recommendations
		WHERE id = ?
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmtOut.Close()

	rec := Recommendation{}

	err = stmtOut.QueryRow(id).Scan(
		&rec.ID, &rec.UserID, &rec.Title, &rec.Type,
		&rec.Body, &rec.Backdrop, &rec.Poster, &rec.Status,
		&rec.CreatedAt, &rec.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &rec, nil
}

// CreateRecommendation creates a new recommendation.
func CreateRecommendation(r *Recommendation, db *sql.DB) (*Recommendation, error) {
	stmtIns, err := db.Prepare(`
		INSERT INTO recommendations (
		user_id, title, type, body, 
		poster, backdrop, status, created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmtIns.Close()

	res, err := stmtIns.Exec(
		r.UserID, r.Title, r.Type, r.Body,
		r.Backdrop, r.Poster, r.Status, r.CreatedAt,
		r.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetRecommendation(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
