package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users", nil)
	if err != nil {
		log.Println(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/users", getUsers).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetUserSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/user/1", nil)
	if err != nil {
		log.Println(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/user/{id}", getUser).Methods("GET")
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
	req, err := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(newUser))
	if err != nil {
		log.Println(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/user", createUser).Methods("POST")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	var upUser = []byte(`{
		"name":"Travis Fox Jr.", 
		"email":"travis_fox@outlook.com",
		"password": "qw12erty"
		}`)
	req, err := http.NewRequest("PUT", "/v1/user/2", bytes.NewBuffer(upUser))
	if err != nil {
		log.Println(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/user/{id}", updateUser).Methods("PUT")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/user/2", nil)
	if err != nil {
		log.Println(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/user/{id}", deleteUser).Methods("DELETE")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
