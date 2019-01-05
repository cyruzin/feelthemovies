package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func TestGetGenresSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/genres", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/genres", getGenres).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetGenreSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/genre/1", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/genre/{id}", getGenre).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

// func TestCreateGenreSuccess(t *testing.T) {

// 	newGenre := struct {
// 		Name string `json:"string"`
// 	}{
// 		"Fantasy",
// 	}

// 	j, err := json.Marshal(newGenre)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	req, err := http.NewRequest("POST", "/v1/genre", bytes.NewBuffer(j))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	rr.Header().Set("Content-Type", "Application/json")

// 	data, err := ioutil.ReadAll(rr.Body)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println(data)

// 	r.HandleFunc("/v1/genre", createGenre).Methods("POST")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// }

// func TestUpdateGenreSuccess(t *testing.T) {

// 	var newGenre = []byte(`{"name":"Music"}`)

// 	req, err := http.NewRequest("PUT", "/v1/genre/2", bytes.NewBuffer(newGenre))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r.HandleFunc("/v1/genre/{id}", updateGenre).Methods("PUT")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// }

// func TestDeleteGenreSuccess(t *testing.T) {
// 	req, err := http.NewRequest("DELETE", "/v1/genre/5", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	data, err := ioutil.ReadAll(rr.Body)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println(data)

// 	r.HandleFunc("/v1/genre/{id}", deleteGenre).Methods("DELETE")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// }
