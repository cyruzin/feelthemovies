package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/app/model"
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

func TestGetRecommendationsPagination(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations?page=2", nil)
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
		t.Error(err)
	}
	req, err := http.NewRequest("POST", "/v1/recommendation", bytes.NewBuffer(ri))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation", createRecommendation).Methods("POST")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}

}

func TestUpdateRecommendation(t *testing.T) {
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
		t.Error(err)
	}
	req, err := http.NewRequest("PUT", "/v1/recommendation/1", bytes.NewBuffer(ri))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation/{id}", updateRecommendation).Methods("PUT")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationFail(t *testing.T) {
	var recItem = []byte(`{"title":"Teste"}`)
	req, err := http.NewRequest("POST", "/v1/recommendation", bytes.NewBuffer(recItem))
	if err != nil {
		t.Error(err)
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
		t.Error(err)
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
