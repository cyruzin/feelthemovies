package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchRecommendationSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_recommendation?query=war", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_recommendation", c.controllers.SearchRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestSearchRecommendationEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_recommendation?query=", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_recommendation", c.controllers.SearchRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestSearchRecommendationMissingField(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_recommendation", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_recommendation", c.controllers.SearchRecommendation)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusBadRequest, status)
	}
}

func TestSearchUserSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_user?query=cyro", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_user", c.controllers.SearchUser)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestSearchGenreSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_genre?query=horror", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_genre", c.controllers.SearchGenre)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestSearchKewordSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_keyword?query=war", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_keyword", c.controllers.SearchKeyword)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestSearchSourceSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/search_source?query=netflix", nil)
	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/search_source", c.controllers.SearchSource)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
