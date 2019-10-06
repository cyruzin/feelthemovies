package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Keyword type is a struct for keywords table.
type Keyword struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// KeywordResult type is a result slice for keywords.
type KeywordResult struct {
	Data *[]Keyword `json:"data"`
}

// GetKeywords retrieves the latest keywords.
func (c *Conn) GetKeywords(ctx context.Context, limit int) (*KeywordResult, error) {
	var result []Keyword

	err := c.db.SelectContext(ctx, &result, queryKeywordsSelect, limit)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &KeywordResult{&result}, nil
}

// GetKeyword retrieves a keyword by ID.
func (c *Conn) GetKeyword(ctx context.Context, id int64) (*Keyword, error) {
	var keyword Keyword

	err := c.db.GetContext(ctx, &keyword, queryKeywordSelectByID, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &keyword, nil
}

// CreateKeyword creates a new keyword.
func (c *Conn) CreateKeyword(ctx context.Context, k *Keyword) error {
	_, err := c.db.ExecContext(ctx, queryKeywordInsert, k.Name, k.CreatedAt, k.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateKeyword updates a keyword by ID.
func (c *Conn) UpdateKeyword(ctx context.Context, id int64, k *Keyword) error {
	result, err := c.db.ExecContext(ctx, queryKeywordUpdate, k.Name, k.UpdatedAt, id)
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

// DeleteKeyword deletes a keyword by ID.
func (c *Conn) DeleteKeyword(ctx context.Context, id int64) error {
	result, err := c.db.ExecContext(ctx, queryKeywordDelete, id)
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
