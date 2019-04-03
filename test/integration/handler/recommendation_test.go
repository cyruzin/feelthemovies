package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

func TestGetRecommendationsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations", nil)

	if err != nil {
		t.Log(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendations", h.h.GetRecommendations)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetRecommendationsPagination(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations?page=2", nil)

	if err != nil {
		t.Log(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendations", h.h.GetRecommendations)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetRecommendationSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", h.h.GetRecommendation)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendation(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	var reqRec struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
	}

	recItem := &model.Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	reqRec.Recommendation = recItem
	reqRec.Genres = []int{3, 5}
	reqRec.Keywords = []int{1, 2}

	ri, err := json.Marshal(reqRec)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/recommendation", bytes.NewBuffer(ri))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/recommendation", h.h.CreateRecommendation)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}

}

func TestUpdateRecommendation(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	var reqRec struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
	}

	recItem := &model.Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Status:    2,
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		UpdatedAt: time.Now(),
	}

	reqRec.Recommendation = recItem
	reqRec.Genres = []int{3, 5}
	reqRec.Keywords = []int{1, 2}

	ri, err := json.Marshal(reqRec)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/recommendation/1", bytes.NewBuffer(ri))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/recommendation/{id}", h.h.UpdateRecommendation)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationFail(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	var recItem = []byte(`{"title":"Teste"}`)

	req, err := http.NewRequest("POST", "/v1/recommendation", bytes.NewBuffer(recItem))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/recommendation", h.h.CreateRecommendation)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationFail(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	var recItem = []byte(`{"title":"Teste"}`)

	req, err := http.NewRequest("PUT", "/v1/recommendation/1", bytes.NewBuffer(recItem))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/recommendation/{id}", h.h.UpdateRecommendation)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationSuccess(t *testing.T) {
	token, err := h.h.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/recommendation/7", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	r.With(h.h.AuthMiddleware).HandleFunc("/v1/recommendation/{id}", h.h.DeleteRecommendation)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
