package model

import (
	"errors"
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

// RecommendationItemResult type is a slice of recommendation items.
type RecommendationItemResult struct {
	Data []*RecommendationItem `json:"data"`
}

// RecommendationItemResponse type is a struct for a final response.
type RecommendationItemResponse struct {
	*RecommendationItem
	Sources []*RecommendationItemSources `json:"sources"`
}

// RecommendationItemSources type is a struct for
// recommendation_item_source pivot table.
type RecommendationItemSources struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationItemFinal type is a struct for
// the final response.
type RecommendationItemFinal struct {
	Data []*RecommendationItemResponse `json:"data"`
}

// RecommendationItemCreate type is a struct for decode
// recommendation item post request.
type RecommendationItemCreate struct {
	*RecommendationItem
	Sources []int  `json:"sources" validate:"required"`
	Year    string `json:"year" validate:"required"`
}

// GetRecommendationItems retrieves all items of
// a given recommendation by ID.
func (db *Conn) GetRecommendationItems(
	id int64,
) (*RecommendationItemResult, error) {
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
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	res := RecommendationItemResult{}
	for rows.Next() {
		rec := RecommendationItem{}
		err = rows.Scan(
			&rec.ID, &rec.RecommendationID, &rec.Name, &rec.TMDBID,
			&rec.Year, &rec.Overview, &rec.Poster, &rec.Backdrop,
			&rec.Trailer, &rec.Commentary, &rec.MediaType, &rec.CreatedAt,
			&rec.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, &rec)
	}
	return &res, nil
}

// GetRecommendationItem retrieves a recommendation
// items by a given ID.
func (db *Conn) GetRecommendationItem(
	id int64,
) (*RecommendationItem, error) {
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
		return nil, err
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
		return nil, err
	}
	return &rec, nil
}

// CreateRecommendationItem creates a new recommendation item.
func (db *Conn) CreateRecommendationItem(
	r *RecommendationItem,
) (*RecommendationItem, error) {
	stmt, err := db.Prepare(`
		INSERT INTO recommendation_items (
		recommendation_id, name, tmdb_id, year, 
		overview, poster, backdrop, trailer, 
		commentary, media_type, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&r.RecommendationID, &r.Name, &r.TMDBID, &r.Year,
		&r.Overview, &r.Poster, &r.Backdrop, &r.Trailer,
		&r.Commentary, &r.MediaType, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	data, err := db.GetRecommendationItem(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateRecommendationItem updates a recommendation
// item by a given ID.
func (db *Conn) UpdateRecommendationItem(
	id int64, r *RecommendationItem,
) (*RecommendationItem, error) {
	stmt, err := db.Prepare(`
		UPDATE recommendation_items
		SET name=?, tmdb_id=?, year=?, overview=?,
		poster=?, backdrop=?, trailer=?, commentary=?,
		media_type=?, updated_at=?
		WHERE id=?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&r.Name, &r.TMDBID, &r.Year, &r.Overview,
		&r.Poster, &r.Backdrop, &r.Trailer, &r.Commentary,
		&r.MediaType, &r.UpdatedAt, &id,
	)
	if err != nil {
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	data, err := db.GetRecommendationItem(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteRecommendationItem deletes a recommendation
// item by a given ID.
func (db *Conn) DeleteRecommendationItem(id int64) error {
	stmt, err := db.Prepare(`
		DELETE 
		FROM recommendation_items
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

// GetRecommendationItemSources retrieves all sources of a given recommendation item.
func (db *Conn) GetRecommendationItemSources(id int64) ([]*RecommendationItemSources, error) {
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

// GetRecommendationItemsTotalRows retrieves the total rows of items of a recommendation.
func (db *Conn) GetRecommendationItemsTotalRows(id int64) (float64, error) {
	stmt, err := db.Prepare(`
		SELECT COUNT(*) 
		FROM recommendation_items 
		WHERE recommendation_id = ?
		`)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	var total float64

	err = stmt.QueryRow(id).Scan(&total)

	if err != nil {
		log.Println(err)
	}

	return total, err
}
