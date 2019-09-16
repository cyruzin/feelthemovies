package helper

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

// Redis expiration time.
const redisTimeout = time.Duration(5 * time.Minute)

// IDParser converts the given ID to int64.
func IDParser(sid string) (int64, error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

// CheckCache checks if the given key exists in cache.
func CheckCache(rc *redis.Client, key string, dest interface{}) (bool, error) {
	cacheValue, _ := rc.Get(key).Result()
	if cacheValue != "" {
		if err := UnmarshalBinary([]byte(cacheValue), dest); err != nil {
			return false, err
		}
	}

	return true, nil
}

// SetCache sets the given key in cache.
func SetCache(rc *redis.Client, key string, dest interface{}) error {
	cacheValue, err := MarshalBinary(dest)
	if err != nil {
		return err
	}

	if err := rc.Set(key, cacheValue, redisTimeout).Err(); err != nil {
		return err
	}

	return nil
}

// HashPassword encrypts a given password using bcrypt algorithm.
func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash checks if the given passwords matches.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MarshalBinary is a implementation of BinaryMarshaler interface.
func MarshalBinary(d interface{}) ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalBinary is a implementation of BinaryUnmarshaler interface.
func UnmarshalBinary(d []byte, v interface{}) error {
	err := json.Unmarshal(d, &v)
	if err != nil {
		return err
	}
	return nil
}

// APIMessage is a struct for generic JSON response.
type APIMessage struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// DecodeError handles API errors.
func DecodeError(
	w http.ResponseWriter,
	r *http.Request,
	l *logger.Logger,
	apiErr string,
	code int,
) {
	l.Errorw(
		apiErr,
		"status", code,
		"method", r.Method,
		"end-point", r.RequestURI,
	) // Loggin before JSON response.
	w.WriteHeader(code) // Setting error code.
	e := &APIMessage{apiErr, code}
	if err := json.NewEncoder(w).Encode(e); err != nil {
		w.Write([]byte("Could not encode the payload"))
		return
	}
}

// APIValidator type is a struct for multiple error messages.
type APIValidator struct {
	Errors []*APIMessage `json:"errors"`
}

// ValidatorMessage handles validation error messages.
func ValidatorMessage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	av := &APIValidator{}
	for _, err := range err.(validator.ValidationErrors) {
		v := &APIMessage{Message: "Check the " + err.Field() + " field"}
		av.Errors = append(av.Errors, v)
	}
	if err := json.NewEncoder(w).Encode(av); err != nil {
		w.Write([]byte("Could not encode the payload"))
		return
	}
}

// SearchValidatorMessage handles search validation errors.
func SearchValidatorMessage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	s := &APIMessage{Message: "The query field is empty"}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		w.Write([]byte("Could not encode the payload"))
		return
	}
}
