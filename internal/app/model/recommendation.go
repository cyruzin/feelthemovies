package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/cyruzin/tome"
)

// Recommendation type is a struct for recommendations table.
type Recommendation struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Type      int       `json:"type"`
	Body      string    `json:"body" validate:"required"`
	Poster    string    `json:"poster" validate:"required"`
	Backdrop  string    `json:"backdrop" validate:"required"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RecommendationResult type is a slice of recommendations.
type RecommendationResult struct {
	Data []*Recommendation `json:"data"`
}

// RecommendationResponse type is a struct for a final response.
type RecommendationResponse struct {
	*Recommendation
	Genres   []*RecommendationGenres   `json:"genres"`
	Keywords []*RecommendationKeywords `json:"keywords"`
}

// RecommendationCreate type is a struct for decode
// recommendation post request.
type RecommendationCreate struct {
	*Recommendation
	Genres   []int `json:"genres"`
	Keywords []int `json:"keywords"`
}

// RecommendationGenres type is a struct for
// genre_recommendation pivot table.
type RecommendationGenres struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationKeywords type is a struct for
// keyword_recommendation pivot table.
type RecommendationKeywords struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationPagination type is a struct for
// paginate recommendations results.
type RecommendationPagination struct {
	Data []*RecommendationResponse `json:"data"`
	*tome.Chapter
}

// GetRecommendations retrieves the latest recommendations.
// o = offset | l = limit
func (c *Conn) GetRecommendations(
	o, l int,
) (*RecommendationResult, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		id, 
		user_id, 
		title, 
		type, 
		body, 
		poster, 
		backdrop, 
		status, 
		created_at, 
		updated_at
		FROM recommendations
		WHERE status = ?
		ORDER BY id DESC
		LIMIT ?,?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(1, o, l)
	if err != nil {
		return nil, err
	}
	res := RecommendationResult{}
	for rows.Next() {
		rec := Recommendation{}
		err = rows.Scan(
			&rec.ID,
			&rec.UserID,
			&rec.Title,
			&rec.Type,
			&rec.Body,
			&rec.Poster,
			&rec.Backdrop,
			&rec.Status,
			&rec.CreatedAt,
			&rec.UpdatedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		res.Data = append(res.Data, &rec)
	}
	return &res, nil
}

// GetRecommendation retrieves a recommendation by a given ID.
func (c *Conn) GetRecommendation(id int64) (*Recommendation, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		id, user_id, title, type, body, poster, 
		backdrop, status, created_at, updated_at
		FROM recommendations
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rec := Recommendation{}
	err = stmt.QueryRow(id).Scan(
		&rec.ID, &rec.UserID, &rec.Title, &rec.Type,
		&rec.Body, &rec.Poster, &rec.Backdrop, &rec.Status,
		&rec.CreatedAt, &rec.UpdatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &rec, err
}

// CreateRecommendation creates a new recommendation.
func (c *Conn) CreateRecommendation(
	r *Recommendation,
) (*Recommendation, error) {
	stmt, err := c.db.Prepare(`
		INSERT INTO recommendations (
		user_id, title, type, body, 
		poster, backdrop, status, created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&r.UserID, &r.Title, &r.Type, &r.Body,
		&r.Poster, &r.Backdrop, &r.Status, &r.CreatedAt,
		&r.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	data, err := c.GetRecommendation(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateRecommendation updates a recommendation by a given ID.
func (c *Conn) UpdateRecommendation(
	id int64, r *Recommendation,
) (*Recommendation, error) {
	stmt, err := c.db.Prepare(`
		UPDATE recommendations
		SET title=?, type=?, body=?, poster=?,
		backdrop=?, status=?, updated_at=?
		WHERE id=?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&r.Title, &r.Type, &r.Body, &r.Poster,
		&r.Backdrop, &r.Status, &r.UpdatedAt, &id,
	)
	if err != nil {
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	data, err := c.GetRecommendation(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteRecommendation deletes a recommendation by a given ID.
func (c *Conn) DeleteRecommendation(id int64) error {
	stmt, err := c.db.Prepare(`
		DELETE 
		FROM recommendations
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

// GetRecommendationGenres retrieves all genres of a given recommendation.
func (c *Conn) GetRecommendationGenres(
	id int64,
) ([]*RecommendationGenres, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		g.id, g.name 
		FROM genres AS g
		JOIN genre_recommendation AS gr ON gr.genre_id = g.id
		JOIN recommendations AS r ON r.id = gr.recommendation_id
		WHERE r.id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	recG := []*RecommendationGenres{}
	for rows.Next() {
		rec := RecommendationGenres{}
		err = rows.Scan(
			&rec.ID, &rec.Name,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		recG = append(recG, &rec)
	}
	return recG, nil
}

// GetRecommendationKeywords retrieves all keywords of a given recommendation.
func (c *Conn) GetRecommendationKeywords(
	id int64,
) ([]*RecommendationKeywords, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		k.id, k.name 
		FROM keywords AS k
		JOIN keyword_recommendation AS kr ON kr.keyword_id = k.id
		JOIN recommendations AS r ON r.id = kr.recommendation_id
		WHERE r.id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	recK := []*RecommendationKeywords{}
	for rows.Next() {
		rec := RecommendationKeywords{}
		err = rows.Scan(
			&rec.ID, &rec.Name,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		recK = append(recK, &rec)
	}
	return recK, nil
}

// GetRecommendationTotalRows retrieves the total rows
// of recommendations table.
func (c *Conn) GetRecommendationTotalRows() (int, error) {
	stmt, err := c.db.Prepare("SELECT COUNT(*) FROM recommendations")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var total int
	err = stmt.QueryRow().Scan(&total)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return total, nil
}

// GetRecommendationsAdmin retrieves the last 10
// recommendations without filter.
func (c *Conn) GetRecommendationsAdmin() (*RecommendationResult, error) {
	stmt, err := c.db.Prepare(`
		SELECT 
		id, 
		user_id, 
		title, 
		type, 
		body, 
		poster, 
		backdrop, 
		status, 
		created_at, 
		updated_at
		FROM recommendations
		ORDER BY id DESC
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(10)
	if err != nil {
		return nil, err
	}
	res := RecommendationResult{}
	for rows.Next() {
		rec := Recommendation{}
		err = rows.Scan(
			&rec.ID,
			&rec.UserID,
			&rec.Title,
			&rec.Type,
			&rec.Body,
			&rec.Poster,
			&rec.Backdrop,
			&rec.Status,
			&rec.CreatedAt,
			&rec.UpdatedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		res.Data = append(res.Data, &rec)
	}
	return &res, nil
}
