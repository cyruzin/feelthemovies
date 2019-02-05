package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

func TestGetRecommendationItemsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_items/1", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation_items/{id}", getRecommendationItems).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetRecommendationItemSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_item/1", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation_item/{id}", getRecommendationItem).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationItemSuccess(t *testing.T) {
	var reqRec struct {
		*model.RecommendationItem
		Sources []int  `json:"sources" validate:"required"`
		Year    string `json:"year" validate:"required"`
	}
	recItem := &model.RecommendationItem{
		Backdrop:         "uhashuas",
		Poster:           "kaoskaos",
		Commentary:       "Foda!",
		Overview:         "uhaushauhs",
		MediaType:        "movie",
		Trailer:          "akska",
		RecommendationID: 1,
		Name:             "John Wick",
		TMDBID:           2312,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	reqRec.RecommendationItem = recItem
	reqRec.Sources = []int{3, 5}
	reqRec.Year = "2017-12-24"
	ri, err := json.Marshal(reqRec)
	if err != nil {
		req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(ri))
		if err != nil {
			t.Error(err)
		}
		rr := httptest.NewRecorder()
		r.HandleFunc("/v1/recommendation_item", createRecommendationItem).Methods("POST")
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
		}
	}
}

func TestUpdateRecommendationItemSuccess(t *testing.T) {
	var reqRec struct {
		*model.RecommendationItem
		Sources []int  `json:"sources" validate:"required"`
		Year    string `json:"year" validate:"required"`
	}
	recItem := &model.RecommendationItem{
		Backdrop:   "uhashuas",
		Poster:     "kaoskaos",
		Commentary: "Foda!",
		Overview:   "kllllalal",
		MediaType:  "movie",
		Trailer:    "akska",
		Name:       "John Wick",
		TMDBID:     2311,
		UpdatedAt:  time.Now(),
	}
	reqRec.RecommendationItem = recItem
	reqRec.Sources = []int{3, 5}
	reqRec.Year = "2017-12-24"
	ri, err := json.Marshal(reqRec)
	if err != nil {
		req, err := http.NewRequest("PUT", "/v1/recommendation_item/1", bytes.NewBuffer(ri))
		if err != nil {
			t.Error(err)
		}
		rr := httptest.NewRecorder()
		r.HandleFunc("/v1/recommendation_item/{id}", updateRecommendationItem).Methods("PUT")
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
		}
	}
}

func TestCreateRecommendationItemFail(t *testing.T) {
	var recItem = []byte(`{"name":"Teste"}`)
	req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(recItem))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation_item", createRecommendationItem).Methods("POST")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationItemFail(t *testing.T) {
	var recItem = []byte(`{"name":"Teste"}`)
	req, err := http.NewRequest("PUT", "/v1/recommendation_item/1", bytes.NewBuffer(recItem))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation_item/{id}", updateRecommendationItem).Methods("PUT")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationItemSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/recommendation_item/7", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	r.HandleFunc("/v1/recommendation_item/{id}", deleteRecommendationItem).Methods("DELETE")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
