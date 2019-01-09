package handlers

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestAuthRoutes(t *testing.T) {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(authMiddleware)

	publicRoutes(r)
	authRoutes(r)
}
