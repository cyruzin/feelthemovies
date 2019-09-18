package model

import (
	"database/sql"
	"errors"
	"time"
)

// RecommendationItem type is a struct for
// recommendation_items table.
type RecommendationItem struct {
	ID               int64     `json:"id"`
	RecommendationID int64     `db:"recommendation_id" json:"recommendation_id" validate:"required,numeric"`
	Name             string    `json:"name" validate:"required"`
	TMDBID           int64     `db:"tmdb_id" json:"tmdb_id" validate:"required,numeric"`
	Year             time.Time `json:"year"`
	Overview         string    `json:"overview" validate:"required"`
	Poster           string    `json:"poster" validate:"required"`
	Backdrop         string    `json:"backdrop" validate:"required"`
	Trailer          string    `json:"trailer"`
	Commentary       string    `json:"commentary"`
	MediaType        string    `db:"media_type" json:"media_type" validate:"required"`
	Sources          string    `json:"sources"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

// RecommendationItemResult type is a slice
// of recommendation items.
type RecommendationItemResult struct {
	Data *[]RecommendationItem `json:"data"`
}

// RecommendationItemSources type is a struct for
// recommendation_item_source pivot table.
type RecommendationItemSources struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RecommendationItemCreate type is a struct for
// decode recommendation item post request.
type RecommendationItemCreate struct {
	*RecommendationItem
	Sources []int  `json:"sources" validate:"required"`
	Year    string `json:"year" validate:"required"`
}

// GetRecommendationItems retrieves all items of
// a given recommendation by ID.
func (c *Conn) GetRecommendationItems(id int64) (*RecommendationItemResult, error) {
	var recommendationItem []RecommendationItem

	err := c.db.Select(&recommendationItem, queryRecommendationItems, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &RecommendationItemResult{&recommendationItem}, nil
}

// GetRecommendationItem retrieves a recommendation
// items by a given ID.
func (c *Conn) GetRecommendationItem(id int64) (*RecommendationItem, error) {
	var recommendationItem RecommendationItem

	err := c.db.Get(&recommendationItem, queryRecommendationItem, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &recommendationItem, nil
}

// CreateRecommendationItem creates a new recommendation item.
func (c *Conn) CreateRecommendationItem(r *RecommendationItem) (int64, error) {
	result, err := c.db.Exec(
		queryRecommendationItemInsert,
		r.RecommendationID,
		r.Name,
		r.TMDBID,
		r.Year,
		r.Overview,
		r.Poster,
		r.Backdrop,
		r.Trailer,
		r.Commentary,
		r.MediaType,
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

// UpdateRecommendationItem updates a recommendation
// item by a given ID.
func (c *Conn) UpdateRecommendationItem(id int64, r *RecommendationItem) (int64, error) {
	result, err := c.db.Exec(
		queryRecommendationItemUpdate,
		r.Name,
		r.TMDBID,
		r.Year,
		r.Overview,
		r.Poster,
		r.Backdrop,
		r.Trailer,
		r.Commentary,
		r.MediaType,
		r.UpdatedAt,
		id,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return 0, errors.New(errResourceNotFound)
	}

	return id, nil
}

// DeleteRecommendationItem deletes a recommendation
// item by a given ID.
func (c *Conn) DeleteRecommendationItem(id int64) error {
	result, err := c.db.Exec(queryRecommendationItemDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New(errResourceNotFound)
	}

	return nil
}

// GetRecommendationItemSources retrieves all
// sources of a given recommendation item.
func (c *Conn) GetRecommendationItemSources(id int64) (*[]RecommendationItemSources, error) {
	var recommendationItemSources []RecommendationItemSources

	err := c.db.Select(&recommendationItemSources, queryRecommendationItemSources, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &recommendationItemSources, nil
}

// GetRecommendationItemsTotalRows retrieves the total rows of items of a recommendation.
func (c *Conn) GetRecommendationItemsTotalRows(id int64) (float64, error) {
	var totalRows float64

	err := c.db.Get(&totalRows, queryRecommendationItemsTotalRows, id)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return totalRows, err
}
