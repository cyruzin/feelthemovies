package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGenresSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/genres", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/genres", h.h.GetGenres)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetGenreSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/genre/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/genre/{id}", h.h.GetGenre)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateGenreSuccess(t *testing.T) {
	var newGenre = []byte(`{"name":"SpongeBob"}`)
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/genre", bytes.NewBuffer(newGenre))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/genre", h.h.CreateGenre)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateGenreSuccess(t *testing.T) {
	var newGenre = []byte(`{"name":"SquidwardTentacles"}`)
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/genre/2", bytes.NewBuffer(newGenre))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/genre/{id}", h.h.UpdateGenre)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteGenreSuccess(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/genre/9", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/genre/{id}", h.h.DeleteGenre)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestMalformedToken(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/genre/7", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/genre/{id}", h.h.DeleteGenre)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusUnauthorized, status)
	}
}

func TestEmptyToken(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/v1/genre/5", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/genre/{id}", h.h.DeleteGenre)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestInvalidToken(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/v1/genre/10", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/genre/{id}", h.h.DeleteGenre)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer a")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusUnauthorized, status)
	}
}
