package model

import (
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/pkg/helper"
)

func TestGetRecommendationItemssModel(t *testing.T) {
	_, err = db.GetRecommendationItems(1)
	if err != nil {
		t.Errorf("GetRecommendationItems error: %s", err)
	}
}

func TestSearchRecommendationItemsTotalRowsModel(t *testing.T) {
	_, err = db.GetRecommendationItemsTotalRows(1)
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
	ri, err := db.CreateRecommendationItem(recItem)
	if err != nil {
		t.Errorf("CreateRecommendationItemError error: %s", err)
	}
	data, err := helper.ToJSON(ri)
	if err != nil {
		t.Errorf("CreateRecommendationItemError - ToJSON error: %s", err)
	}
	t.Log(data)
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
	ri, err := db.UpdateRecommendationItem(1, recItem)
	if err != nil {
		t.Errorf("UpdateRecommendationItemError error: %s", err)
	}
	data, err := helper.ToJSON(ri)
	if err != nil {
		t.Errorf("UpdateRecommendationItemError - ToJSON error: %s", err)
	}
	t.Log(data)
}

func TestGetRecommendationItemModel(t *testing.T) {
	_, err = db.GetRecommendationItems(1)
	if err != nil {
		t.Errorf("GetRecommendationItem error: %s", err)
	}
}

func TestDeleteRecommendationItemsModel(t *testing.T) {
	_, err = db.DeleteRecommendationItem(3)
	if err != nil {
		t.Errorf("DeleteRecommendationItems error: %s", err)
	}
}
