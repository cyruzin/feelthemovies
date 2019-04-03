package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGenresSuccess(t *testing.T) {

	req, err := http.NewRequest("GET", "/v1/genres", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/genres", h.h.GetGenres)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
