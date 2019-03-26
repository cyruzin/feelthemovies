package handler

import (
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/go-redis/redis"
	validator "gopkg.in/go-playground/validator.v9"
)

// Redis expiration time.
const redisTimeout = time.Duration(5 * time.Minute)

// Error messages.
const (
	errDecode       = "Could not decode the body request"
	errParseInt     = "Could not parse to int"
	errParseFloat   = "Could not parse to float"
	errParseDate    = "Could not parse the date"
	errMarhsal      = "Could not marshal the payload"
	errUnmarshal    = "Could not unmarshal the payload"
	errFetch        = "Could not fetch the item"
	errFetchRows    = "Could not fetch the item total rows"
	errKeySet       = "Could not set the key"
	errKeyUnlink    = "Could not unlink the key"
	errCreate       = "Could not create the item"
	errAttach       = "Could not attach the items"
	errSync         = "Could not sync the items"
	errUpdate       = "Could not update the item"
	errDelete       = "Could not delete the item"
	errSearch       = "Could not do the search"
	errPassHash     = "Could not hash the password"
	errAuth         = "Could not authenticate"
	errUnauthorized = "Unauthorized"
	errInvalidToken = "Invalid API Token"
	errQueryField   = "The query field is required"
	errEmptyRec     = "The recommendation is empty or does not exist"
)

// Setup for handlers package.
type Setup struct {
	h  *model.Conn
	rc *redis.Client
	v  *validator.Validate
}

// NewHandler initiates the setup.
func NewHandler(
	m *model.Conn,
	rc *redis.Client,
	v *validator.Validate,
) *Setup {
	return &Setup{m, rc, v}
}
