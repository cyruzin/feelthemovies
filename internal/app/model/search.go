package model

import (
	"database/sql"
)

// Search type is a struct for search queries.
type Search struct {
	Query string `json:"query" validate:"required"`
	Type  int    `json:"type"`
}

// SearchRecommendation search for recommendations.
func (c *Conn) SearchRecommendation(offset, limit int, search string) (*[]Recommendation, error) {
	var recommendations []Recommendation

	err := c.db.Select(
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
func (c *Conn) SearchUser(search string) (*UserResult, error) {
	var users []User

	err := c.db.Select(&users, querySearchUsers, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &UserResult{&users}, nil
}

// SearchGenre search for genres.
func (c *Conn) SearchGenre(search string) (*GenreResult, error) {
	var genres []Genre

	err := c.db.Select(&genres, querySearchGenres, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &GenreResult{&genres}, nil
}

// SearchKeyword search for keywords.
func (c *Conn) SearchKeyword(search string) (*KeywordResult, error) {
	var keywords []Keyword

	err := c.db.Select(&keywords, querySearchKeywords, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &KeywordResult{&keywords}, nil
}

// SearchSource search for sources.
func (c *Conn) SearchSource(search string) (*SourceResult, error) {
	var sources []Source

	err := c.db.Select(&sources, querySearchSources, "%"+search+"%", 20)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &SourceResult{&sources}, nil
}

// GetSearchRecommendationTotalRows retrieves the total
// rows of recommendations table.
func (c *Conn) GetSearchRecommendationTotalRows(search string) (int, error) {
	var totalRows int

	err := c.db.Get(
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
