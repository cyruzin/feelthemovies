package model

import (
	"context"
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
	Sources []int  `json:"sources" validate:"gte=1"`
	Year    string `json:"year" validate:"required"`
}

// GetRecommendationItems retrieves all items of
// a given recommendation by ID.
func (c *Conn) GetRecommendationItems(ctx context.Context, id int64) (*RecommendationItemResult, error) {
	var recommendationItem []RecommendationItem

	err := c.db.SelectContext(ctx, &recommendationItem, queryRecommendationItems, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &RecommendationItemResult{&recommendationItem}, nil
}

// GetRecommendationItem retrieves a recommendation
// items by a given ID.
func (c *Conn) GetRecommendationItem(ctx context.Context, id int64) (*RecommendationItem, error) {
	var recommendationItem RecommendationItem

	err := c.db.GetContext(ctx, &recommendationItem, queryRecommendationItem, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &recommendationItem, nil
}

// CreateRecommendationItem creates a new recommendation item.
func (c *Conn) CreateRecommendationItem(ctx context.Context, r *RecommendationItem) (int64, error) {
	result, err := c.db.ExecContext(
		ctx,
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
func (c *Conn) UpdateRecommendationItem(ctx context.Context, id int64, r *RecommendationItem) error {
	result, err := c.db.ExecContext(
		ctx,
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

// DeleteRecommendationItem deletes a recommendation
// item by a given ID.
func (c *Conn) DeleteRecommendationItem(ctx context.Context, id int64) error {
	result, err := c.db.ExecContext(ctx, queryRecommendationItemDelete, id)
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

// GetRecommendationItemSources retrieves all
// sources of a given recommendation item.
func (c *Conn) GetRecommendationItemSources(ctx context.Context, id int64) (*SourceResult, error) {
	var sources []Source

	err := c.db.SelectContext(ctx, &sources, queryRecommendationItemSources, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &SourceResult{&sources}, nil
}

// GetRecommendationItemsTotalRows retrieves the total rows of items of a recommendation.
func (c *Conn) GetRecommendationItemsTotalRows(ctx context.Context, id int64) (float64, error) {
	var totalRows float64

	err := c.db.GetContext(ctx, &totalRows, queryRecommendationItemsTotalRows, id)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return totalRows, err
}
