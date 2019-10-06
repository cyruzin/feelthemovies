package model

import (
	"context"
	"database/sql"
)

// Search type is a struct for search queries.
type Search struct {
	Query string `json:"query" validate:"required"`
	Type  int    `json:"type"`
}

// SearchRecommendation search for recommendations.
func (c *Conn) SearchRecommendation(ctx context.Context, offset, limit int, search string) (*[]Recommendation, error) {
	var recommendations []Recommendation

	err := c.db.SelectContext(
		ctx,
		&recommendations,
		querySearchRecommendations,
		"%"+search+"%",
		"%"+search+"%",
		"%"+search+"%",
		offset,
		limit,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &recommendations, nil
}

// SearchUser search for users.
func (c *Conn) SearchUser(ctx context.Context, search string) (*UserResult, error) {
	var users []User

	err := c.db.SelectContext(ctx, &users, querySearchUsers, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &UserResult{&users}, nil
}

// SearchGenre search for genres.
func (c *Conn) SearchGenre(ctx context.Context, search string) (*GenreResult, error) {
	var genres []Genre

	err := c.db.SelectContext(ctx, &genres, querySearchGenres, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &GenreResult{&genres}, nil
}

// SearchKeyword search for keywords.
func (c *Conn) SearchKeyword(ctx context.Context, search string) (*KeywordResult, error) {
	var keywords []Keyword

	err := c.db.SelectContext(ctx, &keywords, querySearchKeywords, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &KeywordResult{&keywords}, nil
}

// SearchSource search for sources.
func (c *Conn) SearchSource(ctx context.Context, search string) (*SourceResult, error) {
	var sources []Source

	err := c.db.SelectContext(ctx, &sources, querySearchSources, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &SourceResult{&sources}, nil
}

// GetSearchRecommendationTotalRows retrieves the total
// rows of recommendations table.
func (c *Conn) GetSearchRecommendationTotalRows(ctx context.Context, search string) (int, error) {
	var totalRows int

	err := c.db.GetContext(
		ctx,
		&totalRows,
		querySearchRecommendationsTotalRows,
		"%"+search+"%",
		"%"+search+"%",
		"%"+search+"%",
	)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return totalRows, nil
}
