package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecommendationItemsSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendation_items/1", nil)

	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_item/{id}", getRecommendationItem).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

// func TestCreateRecommendationItemSuccess(t *testing.T) {
// 	const layout = "2006-01-02"
// 	snapshot := "2017-07-25"
// 	yearParsed, err := time.Parse(layout, snapshot)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	reqRec := struct {
// 		RecommendationID int64     `json:"recommendation_id" validate:"required,numeric"`
// 		Name             string    `json:"name" validate:"required"`
// 		TMDBID           int64     `json:"tmdb_id" validate:"required,numeric"`
// 		Year             time.Time `json:"year"`
// 		Overview         string    `json:"overview" validate:"required"`
// 		Poster           string    `json:"poster" validate:"required"`
// 		Backdrop         string    `json:"backdrop" validate:"required"`
// 		Trailer          string    `json:"trailer"`
// 		Commentary       string    `json:"commentary"`
// 		MediaType        string    `json:"media_type" validate:"required"`
// 		CreatedAt        time.Time `json:"created_at"`
// 		UpdatedAt        time.Time `json:"updated_at"`
// 		Sources          []int     `json:"sources" validate:"required"`
// 	}{
// 		1,
// 		"Aquamarine",
// 		5232,
// 		yearParsed,
// 		"I'll kill them all!",
// 		"OqAmzALlnkasQml",
// 		"pKqnkmAqmlasas",
// 		"/iqnonAsnkas",
// 		"Really great movie!",
// 		"movie",
// 		time.Now(),
// 		time.Now(),
// 		[]int{1, 2},
// 	}

// 	ri, err := json.Marshal(&reqRec)

// 	if err != nil {
// 		t.Logf("Could not marshall to Json: %s", err)
// 	}

// 	req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(ri))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r.HandleFunc("/v1/recommendation_item", createRecommendationItem).Methods("POST")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// }

// func TestUpdateRecommendationItemSuccess(t *testing.T) {

// 	var recItem struct {
// 		*model.RecommendationItem
// 		Sources []int  `json:"sources" validate:"required"`
// 		Year    string `json:"year" validate:"required"`
// 	}

// 	recItem.RecommendationID = 2
// 	recItem.Name = "Aquamarine"
// 	recItem.TMDBID = 5232
// 	recItem.Year = "2017-12-24"
// 	recItem.Overview = "I'll kill them all!"
// 	recItem.Poster = "OqAmzALlnkasQml"
// 	recItem.Backdrop = "pKqnkmAqmlasas"
// 	recItem.Trailer = "/iqnonAsnkas"
// 	recItem.Commentary = "Really great movie!"
// 	recItem.MediaType = "serie"
// 	recItem.Sources = []int{4, 6}

// 	ri, err := json.Marshal(recItem)

// 	if err != nil {
// 		t.Logf("Could not marshall to Json: %s", err)
// 	}

// 	req, err := http.NewRequest("POST", "/v1/recommendation_item", bytes.NewBuffer(ri))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	r.HandleFunc("/v1/recommendation_item/{id}", updateRecommendationItem).Methods("POST")

// 	r.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// }

func TestDeleteRecommendationItemSuccess(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/recommendation_item/7", nil)

	if err != nil {
		log.Println(err)
	}

	rr := httptest.NewRecorder()

	r.HandleFunc("/v1/recommendation_item/{id}", deleteRecommendationItem).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
