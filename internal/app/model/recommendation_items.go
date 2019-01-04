package model

import (
	"database/sql"
	"log"
	"time"
)

// RecommendationItem type is a struct for recommendation_items table.
type RecommendationItem struct {
	ID               int64     `json:"id"`
	RecommendationID int64     `json:"recommendation_id" validate:"required,numeric"`
	Name             string    `json:"name" validate:"required"`
	TMDBID           int64     `json:"tmdb_id" validate:"required,numeric"`
	Year             time.Time `json:"year"`
	Overview         string    `json:"overview" validate:"required"`
	Poster           string    `json:"poster" validate:"required"`
	Backdrop         string    `json:"backdrop" validate:"required"`
	Trailer          string    `json:"trailer"`
	Commentary       string    `json:"commentary"`
	MediaType        string    `json:"media_type" validate:"required"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ResultRecommendationItem type is a slice of recommendation items.
type ResultRecommendationItem struct {
	Data []*RecommendationItem `json:"data"`
}

// ResponseRecommendationItem type is a struct for a final response.
type ResponseRecommendationItem struct {
	*RecommendationItem
	Sources []*RecommendationItemSources `json:"sources"`
}

// RecommendationItemSources type is a struct for recommendation_item_source pivot table.
type RecommendationItemSources struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetRecommendationItems retrieves all items of a given recommendation by ID.
func GetRecommendationItems(id int64, db *sql.DB) (*ResultRecommendationItem, error) {

	stmt, err := db.Prepare(`
		SELECT 
		id, recommendation_id, name, tmdb_id, 
		year, overview, poster, backdrop, 
		trailer, commentary, media_type, created_at, 
		updated_at
		FROM recommendation_items
		WHERE recommendation_id = ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	res := ResultRecommendationItem{}

	for rows.Next() {
		rec := RecommendationItem{}

		err = rows.Scan(
			&rec.ID, &rec.RecommendationID, &rec.Name, &rec.TMDBID,
			&rec.Year, &rec.Overview, &rec.Poster, &rec.Backdrop,
			&rec.Trailer, &rec.Commentary, &rec.MediaType, &rec.CreatedAt,
			&rec.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
		}

		res.Data = append(res.Data, &rec)

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
		&r.RecommendationID, &r.Name, &r.TMDBID, &r.Year,
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
		&r.Name, &r.TMDBID, &r.Year, &r.Overview,
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

// GetRecommendationItemSources retrieves all sources of a given recommendation item.
func GetRecommendationItemSources(id int64, db *sql.DB) ([]*RecommendationItemSources, error) {
	stmt, err := db.Prepare(`
		SELECT 
		s.id, s.name 
		FROM sources AS s
		JOIN recommendation_item_source AS ris ON ris.source_id = s.id 	
		JOIN recommendation_items AS ri ON ri.id = ris.recommendation_item_id	
		WHERE ri.id = ?
	`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	recS := []*RecommendationItemSources{}

	for rows.Next() {
		rec := RecommendationItemSources{}

		err = rows.Scan(
			&rec.ID, &rec.Name,
		)

		if err != nil {
			log.Println(err)
		}

		recS = append(recS, &rec)

	}

	return recS, err
}
