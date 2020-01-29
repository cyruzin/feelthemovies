package helper

import (
	"net/http"
	"strings"

	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

	apiValidator := &APIValidator{}

	for _, err := range err.(validator.ValidationErrors) {
		var customError string

		switch err.Tag() {
		case "required":
			customError = "is required"
		case "email":
			customError = "is not valid"
		case "min":
			customError = "minimum length is " + err.Param()
		case "gte":
			customError = "minimum length is " + err.Param()
		}

		message := &APIMessage{
			Message: "The " + strings.ToLower(err.Field()) + " field " + customError,
		}

		apiValidator.Errors = append(apiValidator.Errors, message)
	}

	if err := json.NewEncoder(w).Encode(apiValidator); err != nil {
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
