package handlers

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/go-redis/redis"
)

// Setup ...
type Setup struct {
	h  *model.Conn
	rc *redis.Client
}

// NewHandler ...
func NewHandler(
	m *model.Conn,
	rc *redis.Client) *Setup {
	return &Setup{m, rc}
}

// Redis expiration time.
const redisTimeout = time.Duration(5 * time.Minute)

// Validator instance
var validate *validator.Validate
