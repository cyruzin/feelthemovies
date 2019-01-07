package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSourcesSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/sources", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/sources", getSources).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetSourceSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/source/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/source/{id}", getSource).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreatSourceSuccess(t *testing.T) {

	var newSource = []byte(`{"name":"NewSource"}`)

	req, err := http.NewRequest("POST", "/v1/source", bytes.NewBuffer(newSource))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/source", createSource).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestUpdateSourceSuccess(t *testing.T) {

	var newSource = []byte(`{"name":"UpdateSource"}`)

	req, err := http.NewRequest("PUT", "/v1/source/2", bytes.NewBuffer(newSource))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/source/{id}", updateSource).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteSourceSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/source/7", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/source/{id}", deleteSource).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
