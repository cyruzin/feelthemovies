package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetKeywordsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/keywords", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keywords", h.h.GetKeywords)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetKeywordSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/keyword/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/keyword/{id}", h.h.GetKeyword)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateKeywordSuccess(t *testing.T) {
	var newKeyword = []byte(`{"name":"Tsunami"}`)
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/keyword", bytes.NewBuffer(newKeyword))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/keyword", h.h.CreateKeyword)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateKeywordSuccess(t *testing.T) {
	var newKeyword = []byte(`{"name":"Witness"}`)
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/keyword/2", bytes.NewBuffer(newKeyword))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/keyword/{id}", h.h.UpdateKeyword)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteKeywordSuccess(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/keyword/9", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/keyword/{id}", h.h.DeleteKeyword)

	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
