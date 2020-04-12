package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	model "github.com/cyruzin/feelthemovies/internal/app/models"
)

func TestGetSourcesSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/sources", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/sources", h.handler.GetSources)

	router.ServeHTTP(rr, req)

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

	router.HandleFunc("/v1/source/{id}", h.handler.GetSource)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateSourceSuccess(t *testing.T) {
	var newSource = []byte(`{"name":"BBC Eleven"}`)
	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/source", bytes.NewBuffer(newSource))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/source", h.handler.CreateSource)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateSourceSuccess(t *testing.T) {
	var newSource = []byte(`{"name":"BBC Twelve"}`)
	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/source/2", bytes.NewBuffer(newSource))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/source/{id}", h.handler.UpdateSource)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteSourceSuccess(t *testing.T) {
	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	newSource := model.Source{
		Name:      "Brand New Test Source",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	queryGenreInsert := "INSERT INTO sources (name, created_at, updated_at) VALUES (?, ?, ?)"

	result, err := h.database.Exec(queryGenreInsert, newSource.Name, newSource.CreatedAt, newSource.UpdatedAt)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/source/"+strconv.FormatInt(id, 10), nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/source/{id}", h.handler.DeleteSource)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
