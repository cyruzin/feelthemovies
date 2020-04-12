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

func TestGetRecommendationsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations", nil)
	if err != nil {
		t.Log(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/recommendations", c.controllers.GetRecommendations)

	router.ServeHTTP(rr, req)

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

	router.HandleFunc("/v1/recommendations", c.controllers.GetRecommendations)

	router.ServeHTTP(rr, req)

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

	router.HandleFunc("/v1/recommendation/{id}", c.controllers.GetRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendation(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
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

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation", c.controllers.CreateRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}

}

func TestUpdateRecommendation(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
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

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation/{id}", c.controllers.UpdateRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationFail(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
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

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation", c.controllers.CreateRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationFail(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
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

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation/{id}", c.controllers.UpdateRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationSuccess(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	var request struct {
		*model.Recommendation
		Genres   []int `json:"genres"`
		Keywords []int `json:"keywords"`
	}

	recommendation := &model.Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Status:    0,
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	request.Recommendation = recommendation
	request.Genres = []int{3, 5}
	request.Keywords = []int{1, 2}

	queryRecommendationInsert := `
	    INSERT INTO recommendations (
		user_id, 
		title, 
		type, 
		body, 
		poster, 
		backdrop, 
		status, 
		created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := c.database.Exec(
		queryRecommendationInsert,
		request.UserID,
		request.Title,
		request.Type,
		request.Body,
		request.Poster,
		request.Backdrop,
		request.Status,
		request.CreatedAt,
		request.UpdatedAt,
	)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/recommendation/"+strconv.FormatInt(id, 10), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation/{id}", c.controllers.DeleteRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetRecommendationGenres(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_genres/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/recommendation_genres/{id}", c.controllers.GetRecommendationGenres)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetRecommendationKeywords(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_keywords/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/recommendation_keywords/{id}", c.controllers.GetRecommendationKeywords)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
