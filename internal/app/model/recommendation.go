package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/cyruzin/tome"
)

// Recommendation type is a struct for the recommendations table.
type Recommendation struct {
	ID        int64     `json:"id"`
	UserID    int64     `db:"user_id" json:"user_id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Type      int       `json:"type"`
	Body      string    `json:"body" validate:"required"`
	Poster    string    `json:"poster" validate:"required"`
	Backdrop  string    `json:"backdrop" validate:"required"`
	Status    int       `json:"status"`
	Genres    string    `json:"genres"`
	Keywords  string    `json:"keywords"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// RecommendationResult type is a result slice for recommendations.
type RecommendationResult struct {
	Data *[]Recommendation `json:"data"`
	*tome.Chapter
}

// RecommendationGenres type is a struct for the
// genre_recommendation pivot table.
type RecommendationGenres struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationKeywords type is a struct for the
// keyword_recommendation pivot table.
type RecommendationKeywords struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationCreate type is a struct for decode
// the recommendation post request.
type RecommendationCreate struct {
	*Recommendation
	Genres   []int `json:"genres"`
	Keywords []int `json:"keywords"`
}

// GetRecommendations retrieves the latest recommendations.
// Receives two parameters, the database offset and limit.
func (c *Conn) GetRecommendations(offset, limit int) (*[]Recommendation, error) {
	var result []Recommendation

	err := c.db.Select(&result, queryRecommendationsSelect, offset, limit)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &result, nil
}

// GetRecommendation retrieves a recommendation by ID.
func (c *Conn) GetRecommendation(id int64) (*Recommendation, error) {
	var recommendation Recommendation

	err := c.db.Get(&recommendation, queryRecommendationSelectByID, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &recommendation, nil
}

// CreateRecommendation creates a new recommendation.
func (c *Conn) CreateRecommendation(r *Recommendation) (int64, error) {
	result, err := c.db.Exec(
		queryRecommendationInsert,
		r.UserID,
		r.Title,
		r.Type,
		r.Body,
		r.Poster,
		r.Backdrop,
		r.Status,
		r.CreatedAt,
		r.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateRecommendation updates a recommendation by a given ID.
func (c *Conn) UpdateRecommendation(id int64, r *Recommendation) error {
	result, err := c.db.Exec(
		queryRecommendationUpdate,
		r.Title,
		r.Type,
		r.Body,
		r.Poster,
		r.Backdrop,
		r.Status,
		r.UpdatedAt,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}

// DeleteRecommendation deletes a recommendation by a given ID.
func (c *Conn) DeleteRecommendation(id int64) error {
	result, err := c.db.Exec(queryRecommendationDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}

// GetRecommendationGenres retrieves all genres of a given recommendation.
func (c *Conn) GetRecommendationGenres(id int64) (*GenreResult, error) {
	var genres []Genre

	err := c.db.Select(&genres, queryRecommendationGenres, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &GenreResult{&genres}, nil
}

// GetRecommendationKeywords retrieves all keywords of a given recommendation.
func (c *Conn) GetRecommendationKeywords(id int64) (*KeywordResult, error) {
	var keywords []Keyword

	err := c.db.Select(&keywords, queryRecommendationKeywords, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &KeywordResult{&keywords}, nil
}

// GetRecommendationTotalRows retrieves the total rows
// of recommendations table.
func (c *Conn) GetRecommendationTotalRows() (int, error) {
	var totalRows int

	err := c.db.Get(&totalRows, queryRecommendationsTotalRows)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return totalRows, nil
}

// GetRecommendationsAdmin retrieves the latest recommendations.
// Admin does not have filter for status.
func (c *Conn) GetRecommendationsAdmin() (*RecommendationResult, error) {
	var result []Recommendation

	err := c.db.Select(&result, queryRecommendationsAdminSelect, 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &RecommendationResult{&result, nil}, nil
}
