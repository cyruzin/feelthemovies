package handler

import (
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
	"github.com/go-redis/redis"
	validator "gopkg.in/go-playground/validator.v9"
)

// Redis expiration time.
const redisTimeout = time.Duration(5 * time.Minute)

// Error messages.
const (
	errDecode          = "Could not decode the body request"
	errParseInt        = "Could not parse to int"
	errParseFloat      = "Could not parse to float"
	errParseDate       = "Could not parse the date"
	errMarhsal         = "Could not marshal the payload"
	errUnmarshal       = "Could not unmarshal the payload"
	errFetch           = "Could not fetch the item"
	errFetchRows       = "Could not fetch the item total rows"
	errKeySet          = "Could not set the key"
	errKeyUnlink       = "Could not unlink the key"
	errCreate          = "Could not create the item"
	errAttach          = "Could not attach the items"
	errSync            = "Could not sync the items"
	errUpdate          = "Could not update the item"
	errDelete          = "Could not delete the item"
	errSearch          = "Could not do the search"
	errPassHash        = "Could not hash the password"
	errAuth            = "Could not authenticate"
	errUnauthorized    = "Unauthorized"
	errEmptyToken      = "An API Token is necessary"
	errInvalidToken    = "Invalid API Token"
	errInvalidJWTToken = "Invalid Token"
	errMalformedToken  = "Malformed token"
	errQueryField      = "The query field is required"
	errEmptyRec        = "The recommendation is empty or does not exist"
)

// Setup for handlers package.
type Setup struct {
	h  *model.Conn
	rc *redis.Client
	v  *validator.Validate
	l  *logger.Logger
}

// NewHandler initiates the setup.
func NewHandler(
	m *model.Conn,
	rc *redis.Client,
	v *validator.Validate,
	l *logger.Logger,
) *Setup {
	return &Setup{m, rc, v, l}
}

// CheckCache checks if the given key exists in cache.
func (s *Setup) CheckCache(key string, dest interface{}) (bool, error) {
	cacheValue, _ := s.rc.Get(key).Result()
	if cacheValue != "" {
		if err := helper.UnmarshalBinary([]byte(cacheValue), dest); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

// SetCache sets the given key in cache.
func (s *Setup) SetCache(key string, dest interface{}) error {
	cacheValue, err := helper.MarshalBinary(dest)
	if err != nil {
		return err
	}

	if err := s.rc.Set(key, cacheValue, redisTimeout).Err(); err != nil {
		return err
	}

	return nil
}

// RemoveCache removes the given key from the cache.
func (s *Setup) RemoveCache(key string) error {
	val, _ := s.rc.Get(key).Result()
	if val != "" {
		_, err := s.rc.Unlink(key).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
