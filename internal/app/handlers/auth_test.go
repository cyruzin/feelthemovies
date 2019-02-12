package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthUserSuccess(t *testing.T) {
	var auth = []byte(`{"email":"xorycx@gmail.com", "password":"-%O1r2y3c487-%"}`)
	req, err := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(auth))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/auth", authUser).Methods("POST")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestAuthUserFailure(t *testing.T) {
	var auth = []byte(`{"email":"xorycx@gmail.com", "password":"123456"}`)
	req, err := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(auth))
	if err != nil {
		t.Error(t)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/auth", authUser).Methods("POST")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusUnauthorized, status)
	}
}
