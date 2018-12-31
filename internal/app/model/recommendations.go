package model

import (
	"database/sql"
	"log"
	"time"
)

// TODO: Need to implement keywords and genres sync function.

// Recommendation type is a struct for recommendations table.
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

// ResultRecommendation type is a slice of recommendations.
type ResultRecommendation []*Recommendation

// GetRecommendations retrieves the latest 20 recommendations.
func GetRecommendations(db *sql.DB) (*ResultRecommendation, error) {

	stmt, err := db.Query(`
		SELECT 
		id, user_id, title, type, 
		body, poster, backdrop, status, 
		created_at, updated_at
		FROM recommendations
		ORDER BY id DESC
		LIMIT ?
	`, 20)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res := ResultRecommendation{}

	for stmt.Next() {
		rec := Recommendation{}

		err = stmt.Scan(
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
	stmt, err := db.Prepare(`
		SELECT 
		id, user_id, title, type, body, poster, 
		backdrop, status, created_at, updated_at
		FROM recommendations
		WHERE id = ?
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rec := Recommendation{}

	err = stmt.QueryRow(id).Scan(
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
	stmt, err := db.Prepare(`
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

	defer stmt.Close()

	res, err := stmt.Exec(
		&r.UserID, &r.Title, &r.Type, &r.Body,
		&r.Backdrop, &r.Poster, &r.Status, &r.CreatedAt,
		&r.UpdatedAt,
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

// UpdateRecommendation updates a recommendation by a given ID.
func UpdateRecommendation(id int64, r *Recommendation, db *sql.DB) (*Recommendation, error) {
	stmt, err := db.Prepare(`
		UPDATE recommendations
		SET title=?, type=?, body=?, poster=?,
		backdrop=?, status=?, updated_at=?
		WHERE id=?
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&r.Title, &r.Type, &r.Body, &r.Poster,
		&r.Backdrop, &r.Status, &r.UpdatedAt, &id,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	data, err := GetRecommendation(id, db)

	if err != nil {
		log.Fatal(err)
	}

	return data, nil

}

// DeleteRecommendation deletes a recommendation by a given ID.
func DeleteRecommendation(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM recommendations
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
