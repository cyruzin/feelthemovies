package validation

import (
	"net/http"
	"strings"

	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/go-playground/validator.v9"
)

var json = jsoniter.ConfigFastest

// APIValidator type is a struct for multiple error messages.
type APIValidator struct {
	Errors []*errhandler.APIMessage `json:"errors"`
}

func validationMap(err validator.FieldError) *errhandler.APIMessage {
	errMap := map[string]string{
		"required": "is required",
		"email":    "is not valid",
		"min":      "minimum length is " + err.Param(),
		"gte":      "minimum length is " + err.Param(),
	}

	return &errhandler.APIMessage{
		Message: "The " + strings.ToLower(err.Field()) + " field " + errMap[err.Tag()],
	}
}

// ValidatorMessage handles validation error messages.
func ValidatorMessage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	apiValidator := &APIValidator{}

	for _, err := range err.(validator.ValidationErrors) {
		apiValidator.Errors = append(apiValidator.Errors, validationMap(err))
	}

	if err := json.NewEncoder(w).Encode(apiValidator); err != nil {
		w.Write([]byte("Could not encode the payload"))
		return
	}
}

// SearchValidatorMessage handles search validation errors.
func SearchValidatorMessage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	s := &errhandler.APIMessage{Message: "The query field is empty"}

	if err := json.NewEncoder(w).Encode(s); err != nil {
		w.Write([]byte("Could not encode the payload"))
		return
	}
}
