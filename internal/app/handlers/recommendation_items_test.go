package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecommendationItemsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_items/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_items/{id}", getRecommendationItems).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetRecommendationItemSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_item/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_item/{id}", getRecommendationItem).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationItemFail(t *testing.T) {

	var recItem = []byte(`{"name":"Teste"}`)

	req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(recItem))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_item", createRecommendationItem).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationItemFail(t *testing.T) {

	var recItem = []byte(`{"name":"Teste"}`)

	req, err := http.NewRequest("PUT", "/v1/recommendation_item/1", bytes.NewBuffer(recItem))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_item/{id}", updateRecommendationItem).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationItemSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/recommendation_item/7", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_item/{id}", deleteRecommendationItem).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
