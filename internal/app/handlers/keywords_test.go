package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetKeywordsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/keywords", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keywords", getKeywords).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetKeywordSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/keyword/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keyword/{id}", getKeyword).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateKeywordSuccess(t *testing.T) {

	var newKeyword = []byte(`{"name":"NewKeyWord"}`)

	req, err := http.NewRequest("POST", "/v1/keyword", bytes.NewBuffer(newKeyword))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keyword", createKeyword).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestUpdateKeywordSuccess(t *testing.T) {

	var newGenre = []byte(`{"name":"UpdateKeyword"}`)

	req, err := http.NewRequest("PUT", "/v1/keyword/2", bytes.NewBuffer(newGenre))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keyword/{id}", updateKeyword).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteKeywordSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/keyword/7", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keyword/{id}", deleteKeyword).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
