package handlers

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestMiddleware(t *testing.T) {
	r := mux.NewRouter()

	log := loggingMiddleware(r)
	auth := authMiddleware(r)

	if log == nil {
		t.Error("MiddlewareError")
	}

	if auth == nil {
		t.Error("MiddlewareError")
	}
}
