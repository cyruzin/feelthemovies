package errhandler

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
)

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
	)

	w.WriteHeader(code)

	e := &APIMessage{apiErr, code}
	if err := json.NewEncoder(w).Encode(e); err != nil {
		w.Write([]byte("Could not encode the payload"))
		return
	}
}
