package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthUserSuccess(t *testing.T) {
	var auth = []byte(`{"email":"admin@admin.com", "password":"password"}`)

	req, err := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(auth))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/auth", h.h.AuthUser)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestAuthUserFailure(t *testing.T) {
	var auth = []byte(`{"email":"admin@admin.com", "password":"123456"}`)

	req, err := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(auth))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/auth", h.h.AuthUser)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusUnauthorized, status)
	}
}
