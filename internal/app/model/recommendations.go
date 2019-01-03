package model

import (
	"database/sql"
	"log"
	"time"
)

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
	Genres    []int     `json:"genres,omitempty"`
	Keywords  []int     `json:"keywords,omitempty"`
}

// ResultRecommendation type is a slice of recommendations.
type ResultRecommendation struct {
	Data []*Recommendation `json:"data"`
}

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
		log.Println(err)
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
			log.Println(err)
		}

		res.Data = append(res.Data, &rec)

	}

	return &res, err
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
		log.Println(err)
	}

	defer stmt.Close()

	rec := Recommendation{}

	err = stmt.QueryRow(id).Scan(
		&rec.ID, &rec.UserID, &rec.Title, &rec.Type,
		&rec.Body, &rec.Backdrop, &rec.Poster, &rec.Status,
		&rec.CreatedAt, &rec.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	return &rec, err
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
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&r.UserID, &r.Title, &r.Type, &r.Body,
		&r.Backdrop, &r.Poster, &r.Status, &r.CreatedAt,
		&r.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	data, err := GetRecommendation(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
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
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&r.Title, &r.Type, &r.Body, &r.Poster,
		&r.Backdrop, &r.Status, &r.UpdatedAt, &id,
	)

	if err != nil {
		log.Println(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	data, err := GetRecommendation(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// DeleteRecommendation deletes a recommendation by a given ID.
func DeleteRecommendation(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM recommendations
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
