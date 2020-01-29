package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	validator "gopkg.in/go-playground/validator.v9"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Redis expiration time.
const redisTimeout = time.Duration(5 * time.Minute)

// Error messages.
const (
	errDecode          = "Could not decode the body request"
	errParseInt        = "Could not parse to int"
	errParseFloat      = "Could not parse to float"
	errParseDate       = "Could not parse the date"
	errMarshal         = "Could not marshal the payload"
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
	errEmptyToken      = "JWT Token not provided"
	errInvalidJWTToken = "Invalid JWT Token"
	errMalformedToken  = "Malformed token"
	errGeneratingToken = "Could not generate the Token"
	errQueryField      = "The query field is required"
	errEmptyRec        = "The recommendation is empty, add at least one item"
)

// Setup for handlers package.
type Setup struct {
	model     *model.Conn
	redis     *redis.Client
	validator *validator.Validate
	logger    *logger.Logger
}

// NewHandler initiates the setup.
func NewHandler(
	model *model.Conn,
	redis *redis.Client,
	validator *validator.Validate,
	logger *logger.Logger,
) *Setup {
	return &Setup{model, redis, validator, logger}
}

// CheckCache checks if the given key exists in cache.
func (s *Setup) CheckCache(ctx context.Context, key string, dest interface{}) (bool, error) {
	cacheValue, _ := s.redis.WithContext(ctx).Get(key).Result()
	if cacheValue != "" {
		if err := helper.UnmarshalBinary([]byte(cacheValue), dest); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

// SetCache sets the given key in cache.
func (s *Setup) SetCache(ctx context.Context, key string, dest interface{}) error {
	cacheValue, err := helper.MarshalBinary(dest)
	if err != nil {
		return err
	}

	if err := s.redis.WithContext(ctx).Set(key, cacheValue, redisTimeout).Err(); err != nil {
		return err
	}

	return nil
}

// RemoveCache removes the given key from the cache.
func (s *Setup) RemoveCache(ctx context.Context, key string) error {
	val, _ := s.redis.WithContext(ctx).Get(key).Result()
	if val != "" {
		_, err := s.redis.Unlink(key).Result()
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

// GenerateCacheKey generates a cache key base on
// the given key and returns a string.
func (s *Setup) GenerateCacheKey(params url.Values, key string) string {
	var cacheKey string

	if params["page"] != nil && params["query"] == nil { // Only page param.
		cacheKey = fmt.Sprintf("%s?page=%s", key, params["page"][0])
	} else if params["page"] == nil && params["query"] != nil { // Only query param.
		cacheKey = fmt.Sprintf("%s?query=%s", key, params["query"][0])
	} else if params["page"] != nil && params["query"] != nil { // Both.
		cacheKey = fmt.Sprintf("?query=%s&page=%s", params["query"][0], params["page"][0])
	} else {
		cacheKey = key
	}

	return cacheKey
}

// IDParser converts the given ID to int64.
func (s *Setup) IDParser(sid string) (int64, error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

// PageParser checks if page string exists in the
// given URL params. If exists, it will be parsed to
// int and returned. If some error occurs, the default
// value will be returned.
//
// Default value: 1.
func (s *Setup) PageParser(params url.Values) (int, error) {
	newPage := 1
	if params["page"] != nil && params["page"][0] != "" {
		newPage, err := strconv.Atoi(params["page"][0])
		if err != nil {
			return newPage, err
		}
		return newPage, nil
	}

	return newPage, nil
}

// ToJSON returns a JSON response.
func (s *Setup) ToJSON(
	w http.ResponseWriter,
	httpStatus int,
	dest interface{},
) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(dest)
}
