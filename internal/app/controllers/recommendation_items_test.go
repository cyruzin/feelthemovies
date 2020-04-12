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

func TestGetRecommendationItemsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_items/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/recommendation_items/{id}", c.controllers.GetRecommendationItems)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetRecommendationItemSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_item/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/recommendation_item/{id}", c.controllers.GetRecommendationItem)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationItemSuccess(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

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
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(ri))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation_item", c.controllers.CreateRecommendationItem)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}

}

func TestUpdateRecommendationItemSuccess(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	var reqRec struct {
		*model.RecommendationItem
		Sources []int  `json:"sources" validate:"required"`
		Year    string `json:"year" validate:"required"`
	}

	recItem := &model.RecommendationItem{
		Backdrop:         "uhashuas",
		Poster:           "kaoskaos",
		Commentary:       "Foda!",
		Overview:         "kllllalal",
		MediaType:        "movie",
		Trailer:          "akska",
		RecommendationID: 1,
		Name:             "John Wick: Chapter 2",
		TMDBID:           2311,
		UpdatedAt:        time.Now(),
	}

	reqRec.RecommendationItem = recItem
	reqRec.Sources = []int{3, 5}
	reqRec.Year = "2017-12-24"

	ri, err := json.Marshal(reqRec)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/recommendation_item/1", bytes.NewBuffer(ri))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation_item/{id}", c.controllers.UpdateRecommendationItem)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateRecommendationItemFail(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	var recItem = []byte(`{"name":"Teste"}`)

	req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(recItem))
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation_item", c.controllers.CreateRecommendationItem)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestUpdateRecommendationItemFail(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	var recItem = []byte(`{"name":"Teste"}`)

	req, err := http.NewRequest("PUT", "/v1/recommendation_item/1", bytes.NewBuffer(recItem))
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation_item/{id}", c.controllers.UpdateRecommendationItem)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestDeleteRecommendationItemSuccess(t *testing.T) {
	token, err := c.controllers.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	var request struct {
		*model.RecommendationItem
		Sources []int  `json:"sources" validate:"required"`
		Year    string `json:"year" validate:"required"`
	}

	item := &model.RecommendationItem{
		Backdrop:         "uhashuas",
		Poster:           "kaoskaos",
		Commentary:       "Foda!",
		Overview:         "kllllalal",
		MediaType:        "movie",
		Trailer:          "akska",
		RecommendationID: 1,
		Name:             "John Wick: Chapter 2",
		TMDBID:           2311,
		UpdatedAt:        time.Now(),
	}

	request.RecommendationItem = item
	request.Sources = []int{3, 5}
	request.Year = "2017-12-24"
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	queryRecommendationItemInsert := `
		INSERT INTO recommendation_items (
		recommendation_id, 
		name, 
		tmdb_id, 
		year, 
		overview, 
		poster, 
		backdrop, 
		trailer, 
		commentary, 
		media_type, 
		created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := c.database.Exec(
		queryRecommendationItemInsert,
		request.RecommendationID,
		request.Name,
		request.TMDBID,
		request.Year,
		request.Overview,
		request.Poster,
		request.Backdrop,
		request.Trailer,
		request.Commentary,
		request.MediaType,
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

	req, err := http.NewRequest("DELETE", "/v1/recommendation_item/"+strconv.FormatInt(id, 10), nil)
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(c.controllers.AuthMiddleware).HandleFunc("/v1/recommendation_item/{id}", c.controllers.DeleteRecommendationItem)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestGetRecommendationItemSources(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_item_sources/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc(
		"/v1/recommendation_item_sources/{id}",
		c.controllers.GetRecommendationItemSources,
	)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
