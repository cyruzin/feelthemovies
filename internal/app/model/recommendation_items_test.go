package model

import (
	"testing"
	"time"
)

func TestGetRecommendationItemssModel(t *testing.T) {
	_, err := db.GetRecommendationItems(1)
	if err != nil {
		t.Errorf("GetRecommendationItems error: %s", err)
	}
}

func TestSearchRecommendationItemsTotalRowsModel(t *testing.T) {
	_, err := db.GetRecommendationItemsTotalRows(1)
	if err != nil {
		t.Errorf("RecommendationItemsTotalRows error: %s", err)
	}
}

func TestCreateRecommendationItemsModel(t *testing.T) {
	year := "2017-12-24"
	yearParsed, err := time.Parse("2006-01-02", year)
	if err != nil {
		t.Errorf("CreateRecommendationItemError - Time Parse -error: %s", err)
	}
	recItem := &RecommendationItem{
		RecommendationID: 1,
		Name:             "Wonder Woman",
		TMDBID:           2321,
		Year:             yearParsed,
		Overview:         "Lorem ipsum... Wonder Woman",
		Poster:           "hoasmoajs",
		Backdrop:         "huPakslams",
		Trailer:          "/heuapsmasl",
		Commentary:       "Amazing movie!",
		MediaType:        "movie",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	_, err = db.CreateRecommendationItem(recItem)
	if err != nil {
		t.Errorf("CreateRecommendationItemError error: %s", err)
	}
}

func TestUpdateRecommendationItemsModel(t *testing.T) {
	year := "2017-12-24"
	yearParsed, err := time.Parse("2006-01-02", year)
	if err != nil {
		t.Errorf("UpdateRecommendationItemError - Time Parse - error: %s", err)
	}
	recItem := &RecommendationItem{
		Name:       "Wonder Woman",
		TMDBID:     2321,
		Year:       yearParsed,
		Overview:   "Lorem ipsum... Wonder Woman",
		Poster:     "hoasmoajs",
		Backdrop:   "huPakslams",
		Trailer:    "/heuapsmasl",
		Commentary: "Amazing movie!",
		MediaType:  "movie",
		UpdatedAt:  time.Now(),
	}
	_, err = db.UpdateRecommendationItem(1, recItem)
	if err != nil {
		t.Errorf("UpdateRecommendationItemError error: %s", err)
	}
}

func TestGetRecommendationItemModel(t *testing.T) {
	_, err := db.GetRecommendationItems(1)
	if err != nil {
		t.Errorf("GetRecommendationItem error: %s", err)
	}
}

func TestDeleteRecommendationItemsModel(t *testing.T) {
	if err := db.DeleteRecommendationItem(3); err != nil {
		t.Errorf("DeleteRecommendationItems error: %s", err)
	}
}
