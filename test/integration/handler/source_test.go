package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSourcesSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/sources", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/sources", h.h.GetSources)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetSourceSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/source/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/source/{id}", h.h.GetSource)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateSourceSuccess(t *testing.T) {
	var newSource = []byte(`{"name":"BBC Eleven"}`)
	token, err := h.h.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/source", bytes.NewBuffer(newSource))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/source", h.h.CreateSource)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateSourceSuccess(t *testing.T) {
	var newSource = []byte(`{"name":"BBC Twelve"}`)
	token, err := h.h.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/source/2", bytes.NewBuffer(newSource))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/source/{id}", h.h.UpdateSource)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteSourceSuccess(t *testing.T) {
	token, err := h.h.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/source/7", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/source/{id}", h.h.DeleteSource)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
