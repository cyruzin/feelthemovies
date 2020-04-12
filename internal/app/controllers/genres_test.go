package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	model "github.com/cyruzin/feelthemovies/internal/app/models"
	"github.com/go-chi/chi"
)

func TestGetGenresSuccess(t *testing.T) {
	router := chi.NewRouter()
	req, err := http.NewRequest("GET", "/v1/genres", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/genres", c.controllers.GetGenres)

	router.ServeHTTP(rr, req)

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

	router.HandleFunc("/v1/genre/{id}", c.controllers.GetGenre)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateGenreSuccess(t *testing.T) {
	var newGenre = []byte(`{"name":"SpongeBob"}`)

	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/genre", bytes.NewBuffer(newGenre))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/genre", c.controllers.CreateGenre)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateGenreSuccess(t *testing.T) {
	var newGenre = []byte(`{"name":"SquidwardTentacles"}`)
	token, err := c.controllers.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/genre/2", bytes.NewBuffer(newGenre))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/genre/{id}", c.controllers.UpdateGenre)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteGenreSuccess(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	newGenre := model.Genre{
		Name:      "Brand New Test Genre",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	queryGenreInsert := "INSERT INTO genres (name, created_at, updated_at) VALUES (?, ?, ?)"

	result, err := c.database.Exec(queryGenreInsert, newGenre.Name, newGenre.CreatedAt, newGenre.UpdatedAt)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/genre/"+strconv.FormatInt(id, 10), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/genre/{id}", c.controllers.DeleteGenre)

	req.Header.Add("Authorization", "Bearer "+token)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestMalformedToken(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/genre/9", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/genre/{id}", c.controllers.DeleteGenre)

	req.Header.Add("Authorization", token)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusUnauthorized, status)
	}
}

func TestEmptyToken(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/genre/9", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/genre/{id}", c.controllers.DeleteGenre)

	req.Header.Add("Authorization", "")

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestInvalidToken(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/genre/9", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/genre/{id}", c.controllers.DeleteGenre)

	req.Header.Add("Authorization", "Bearer a")

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusUnauthorized, status)
	}
}
