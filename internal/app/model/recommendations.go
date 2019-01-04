package model

import (
	"database/sql"
	"log"
	"time"
)

// Recommendation type is a struct for recommendations table.
type Recommendation struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id" validate:"required,numeric"`
	Title     string    `json:"title" validate:"required"`
	Type      int       `json:"type" validate:"required,numeric"`
	Body      string    `json:"body" validate:"required"`
	Poster    string    `json:"poster" validate:"required"`
	Backdrop  string    `json:"backdrop" validate:"required"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ResultRecommendation type is a slice of recommendations.
type ResultRecommendation struct {
	Data []*Recommendation `json:"data"`
}

// ResponseRecommendation type is a struct for a final response.
type ResponseRecommendation struct {
	*Recommendation
	Genres   []*RecommendationGenres   `json:"genres"`
	Keywords []*RecommendationKeywords `json:"keywords"`
}

// RecommendationGenres type is a struct for genre_recommendation pivot table.
type RecommendationGenres struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationKeywords type is a struct for keyword_recommendation pivot table.
type RecommendationKeywords struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetRecommendations retrieves the latest 20 recommendations.
func GetRecommendations(db *sql.DB) (*ResultRecommendation, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, user_id, title, type, 
		body, poster, backdrop, status, 
		created_at, updated_at
		FROM recommendations
		ORDER BY id DESC
		LIMIT ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(10)

	res := ResultRecommendation{}

	for rows.Next() {
		rec := Recommendation{}

		err = rows.Scan(
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

// GetRecommendationGenres retrieves all genres of a given recommendation.
func GetRecommendationGenres(id int64, db *sql.DB) ([]*RecommendationGenres, error) {
	stmt, err := db.Prepare(`
		SELECT 
		g.id, g.name 
		FROM genres AS g
		JOIN genre_recommendation AS gr ON gr.genre_id = g.id
		JOIN recommendations AS r ON r.id = gr.recommendation_id
		WHERE r.id = ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	recG := []*RecommendationGenres{}

	for rows.Next() {
		rec := RecommendationGenres{}

		err = rows.Scan(
			&rec.ID, &rec.Name,
		)

		if err != nil {
			log.Println(err)
		}

		recG = append(recG, &rec)

	}

	return recG, err
}

// GetRecommendationKeywords retrieves all keywords of a given recommendation.
func GetRecommendationKeywords(id int64, db *sql.DB) ([]*RecommendationKeywords, error) {
	stmt, err := db.Prepare(`
		SELECT 
		k.id, k.name 
		FROM keywords AS k
		JOIN keyword_recommendation AS kr ON kr.keyword_id = k.id
		JOIN recommendations AS r ON r.id = kr.recommendation_id
		WHERE r.id = ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	recK := []*RecommendationKeywords{}

	for rows.Next() {
		rec := RecommendationKeywords{}

		err = rows.Scan(
			&rec.ID, &rec.Name,
		)

		if err != nil {
			log.Println(err)
		}

		recK = append(recK, &rec)

	}

	return recK, err
}
