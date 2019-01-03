package model

import (
	"database/sql"
	"log"
	"time"
)

// RecommendationItem type is a struct for recommendation_items table.
type RecommendationItem struct {
	ID               int64     `json:"id"`
	RecommendationID int64     `json:"recommendation_id"`
	Name             string    `json:"name"`
	TMDBID           int64     `json:"tmdb_id"`
	Year             string    `json:"year"`
	YearParsed       time.Time `json:"-"`
	Overview         string    `json:"overview"`
	Poster           string    `json:"poster"`
	Backdrop         string    `json:"backdrop"`
	Trailer          string    `json:"trailer"`
	Commentary       string    `json:"commentary"`
	MediaType        string    `json:"media_type"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ResultRecommendationItem type is a slice of recommendation items.
type ResultRecommendationItem []*RecommendationItem

// GetRecommendationItems retrieves all items of a given recommendation by ID.
func GetRecommendationItems(id int64, db *sql.DB) (*ResultRecommendationItem, error) {

	stmt, err := db.Query(`
		SELECT 
		id, recommendation_id, name, tmdb_id, 
		year, overview, poster, backdrop, 
		trailer, commentary, media_type, created_at, 
		updated_at
		FROM recommendation_items
		WHERE recommendation_id = ?
	`, id)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res := ResultRecommendationItem{}

	for stmt.Next() {
		rec := RecommendationItem{}

		err = stmt.Scan(
			&rec.ID, &rec.RecommendationID, &rec.Name, &rec.TMDBID,
			&rec.Year, &rec.Overview, &rec.Poster, &rec.Backdrop,
			&rec.Trailer, &rec.Commentary, &rec.MediaType, &rec.CreatedAt,
			&rec.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res = append(res, &rec)

	}

	return &res, err
}

// GetRecommendationItem retrieves a recommendation items by a given ID.
func GetRecommendationItem(id int64, db *sql.DB) (*RecommendationItem, error) {
	stmt, err := db.Prepare(`
		SELECT 
		id, recommendation_id, name, tmdb_id, 
		year, overview, poster, backdrop, 
		trailer, commentary, media_type, created_at, 
		updated_at
		FROM recommendation_items
		WHERE id = ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rec := RecommendationItem{}

	err = stmt.QueryRow(id).Scan(
		&rec.ID, &rec.RecommendationID, &rec.Name, &rec.TMDBID,
		&rec.Year, &rec.Overview, &rec.Poster, &rec.Backdrop,
		&rec.Trailer, &rec.Commentary, &rec.MediaType, &rec.CreatedAt,
		&rec.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	return &rec, err
}

// CreateRecommendationItem creates a new recommendation item.
func CreateRecommendationItem(r *RecommendationItem, db *sql.DB) (*RecommendationItem, error) {
	stmt, err := db.Prepare(`
		INSERT INTO recommendation_items (
		recommendation_id, name, tmdb_id, year, 
		overview, poster, backdrop, trailer, 
		commentary, media_type, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&r.RecommendationID, &r.Name, &r.TMDBID, &r.YearParsed,
		&r.Overview, &r.Poster, &r.Backdrop, &r.Trailer,
		&r.Commentary, &r.MediaType, &r.CreatedAt, &r.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	data, err := GetRecommendationItem(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// UpdateRecommendationItem updates a recommendation item by a given ID.
func UpdateRecommendationItem(id int64, r *RecommendationItem, db *sql.DB) (*RecommendationItem, error) {
	stmt, err := db.Prepare(`
		UPDATE recommendation_items
		SET name=?, tmdb_id=?, year=?, overview=?,
		poster=?, backdrop=?, trailer=?, commentary=?,
		media_type=?, updated_at=?
		WHERE id=?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&r.Name, &r.TMDBID, &r.YearParsed, &r.Overview,
		&r.Poster, &r.Backdrop, &r.Trailer, &r.Commentary,
		&r.MediaType, &r.UpdatedAt, &id,
	)

	if err != nil {
		log.Println(err)
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	data, err := GetRecommendationItem(id, db)

	if err != nil {
		log.Println(err)
	}

	return data, err

}

// DeleteRecommendationItem deletes a recommendation item by a given ID.
func DeleteRecommendationItem(id int64, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(`
		DELETE 
		FROM recommendation_items
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
