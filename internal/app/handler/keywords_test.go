package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

func TestGetKeywordsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/keywords", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/keywords", h.handler.GetKeywords)

	router.ServeHTTP(rr, req)

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

	router.HandleFunc("/v1/keyword/{id}", h.handler.GetKeyword)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateKeywordSuccess(t *testing.T) {
	var newKeyword = []byte(`{"name":"Tsunami"}`)
	token, err := h.handler.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/keyword", bytes.NewBuffer(newKeyword))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/keyword", h.handler.CreateKeyword)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateKeywordSuccess(t *testing.T) {
	var newKeyword = []byte(`{"name":"Witness"}`)
	token, err := h.handler.GenerateToken(info)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/keyword/2", bytes.NewBuffer(newKeyword))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/keyword/{id}", h.handler.UpdateKeyword)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteKeywordSuccess(t *testing.T) {
	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	newKeyword := model.Keyword{
		Name:      "Brand New Test Keyword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	queryGenreInsert := "INSERT INTO keywords (name, created_at, updated_at) VALUES (?, ?, ?)"

	result, err := h.database.Exec(queryGenreInsert, newKeyword.Name, newKeyword.CreatedAt, newKeyword.UpdatedAt)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/keyword/"+strconv.FormatInt(id, 10), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/keyword/{id}", h.handler.DeleteKeyword)

	req.Header.Add("Authorization", "Bearer "+token)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
