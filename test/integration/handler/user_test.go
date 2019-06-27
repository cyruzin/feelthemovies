package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/users", h.h.GetUsers)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetUserSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/user/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/user/{id}", h.h.GetUser)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateUserSuccess(t *testing.T) {
	var newUser = []byte(`{
		"name":"Travis Fox", 
		"email":"travis_fox@outlook.com",
		"password": "qw12erty"
		}`)
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(newUser))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/user", h.h.CreateUser)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	var upUser = []byte(`{
		"name":"Travis Fox Jr.", 
		"email":"travis_fox_jr@outlook.com",
		"password": "qw12erty"
		}`)
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/user/1", bytes.NewBuffer(upUser))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/user/{id}", h.h.UpdateUser)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/user/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/user/{id}", h.h.DeleteUser)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
