package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareBlock(t *testing.T) {

	req, err := http.NewRequest("GET", "/v1/recommendations", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.Use(authMiddleware)

	r.HandleFunc("/v1/recommendations", getRecommendations).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusForbidden, status)
	}
}

func TestMiddlewareSuccess(t *testing.T) {

	req, err := http.NewRequest("GET", "/v1/recommendations", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()
	rr.Header().Set("Content-Type", "Application/json")
	rr.Header().Set("Api-Token", "bb9b6ed1-8688-44f4-9f35-f75e62ef83f1")

	r.HandleFunc("/v1/recommendations", getRecommendations).Methods("GET")
	r.ServeHTTP(rr, req)
	r.Use(authMiddleware)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
