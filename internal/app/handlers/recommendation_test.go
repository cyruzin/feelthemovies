package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecommendationsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendations", getRecommendations).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetRecommendationSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", getRecommendation).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationFail(t *testing.T) {

	var recItem = []byte(`{"title":"Teste"}`)

	req, err := http.NewRequest("POST", "/v1/recommendation", bytes.NewBuffer(recItem))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation", createRecommendation).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationFail(t *testing.T) {

	var recItem = []byte(`{"title":"Teste"}`)

	req, err := http.NewRequest("PUT", "/v1/recommendation/1", bytes.NewBuffer(recItem))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", updateRecommendation).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/recommendation/7", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", deleteRecommendation).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
