package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestGetRecommendationsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendations", getRecommendations).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetRecommendationSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", getRecommendation).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendation(t *testing.T) {

	newRec := &model.Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	validate = validator.New()
	err = validate.Struct(newRec)

	if err != nil {
		t.Errorf("CreateRecommendation - Validation - error: %s", err)
	}

	rec, err := db.CreateRecommendation(newRec)

	if err != nil {
		t.Errorf("CreateRecommendation error: %s", err)
	}

	keywords := make(map[int64][]int)
	genres := make(map[int64][]int)

	keywords[rec.ID] = []int{1, 2}
	genres[rec.ID] = []int{3, 5}

	_, err = helper.Attach(keywords, "keyword_recommendation", db.DB)

	if err != nil {
		t.Errorf("CreateRecommendation - Attach - error: %s", err)
	}

	_, err = helper.Attach(genres, "genre_recommendation", db.DB)

	if err != nil {
		t.Errorf("CreateRecommendation - Attach - error: %s", err)
	}
}

func TestUpdateRecommendation(t *testing.T) {

	// Check status
	validate = validator.New()
	err = validate.Var(1, "required,min=1,max=2")

	if err != nil {
		t.Errorf("UpdateRecommendation - Validation - error: %s", err)
	}

	upRec := &model.Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Status:    2,
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		UpdatedAt: time.Now(),
	}

	validate = validator.New()
	err = validate.Struct(upRec)

	if err != nil {
		t.Errorf("UpdateRecommendation - Validation - error: %s", err)
	}

	rec, err := db.UpdateRecommendation(2, upRec)

	if err != nil {
		t.Errorf("UpdateRecommendation error: %s", err)
	}

	keywords := make(map[int64][]int)
	genres := make(map[int64][]int)

	keywords[rec.ID] = []int{1, 2}
	genres[rec.ID] = []int{3, 5}

	_, err = helper.Attach(keywords, "keyword_recommendation", db.DB)

	if err != nil {
		t.Errorf("UpdateRecommendation - Attach - error: %s", err)
	}

	_, err = helper.Attach(genres, "genre_recommendation", db.DB)

	if err != nil {
		t.Errorf("UpdateRecommendation - Attach - error: %s", err)
	}
}

func TestCreateRecommendationFail(t *testing.T) {

	var recItem = []byte(`{"title":"Teste"}`)

	req, err := http.NewRequest("POST", "/v1/recommendation", bytes.NewBuffer(recItem))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation", createRecommendation).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationFail(t *testing.T) {

	var recItem = []byte(`{"title":"Teste"}`)

	req, err := http.NewRequest("PUT", "/v1/recommendation/1", bytes.NewBuffer(recItem))

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", updateRecommendation).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/recommendation/7", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation/{id}", deleteRecommendation).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
