package model

import (
	"testing"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

var db, err = Connect()

func TestSearchRecommendationModel(t *testing.T) {
	_, err = db.SearchRecommendation(0, 10, "war")
	if err != nil {
		t.Errorf("SearchRecommendations error: %s", err)
	}
}

func TestSearchRecommendationTotalRowsModel(t *testing.T) {
	_, err = db.GetSearchRecommendationTotalRows("war")
	if err != nil {
		t.Errorf("SearchRecommendations error: %s", err)
	}
}

func TestSearchUserModel(t *testing.T) {
	user, err := db.SearchUser("Cyro")
	if err != nil {
		t.Errorf("SearchUser error: %s", err)
	}
	data, err := helper.ToJSON(user)
	if err != nil {
		t.Errorf("SearchUser - ToJSON - error: %s", err)
	}
	t.Log(data)
}

func TestSearchGenreModel(t *testing.T) {
	genre, err := db.SearchGenre("Comedy")
	if err != nil {
		t.Errorf("SearchGenre error: %s", err)
	}
	data, err := helper.ToJSON(genre)
	if err != nil {
		t.Errorf("SearchGenre - ToJSON - error: %s", err)
	}
	t.Log(data)
}

func TestSearchKeywordModel(t *testing.T) {
	keyword, err := db.SearchKeyword("Shark")
	if err != nil {
		t.Errorf("SearchKeyword error: %s", err)
	}
	data, err := helper.ToJSON(keyword)
	if err != nil {
		t.Errorf("SearchKeyword - ToJSON - error: %s", err)
	}
	t.Log(data)
}

func TestSearchSourceModel(t *testing.T) {
	source, err := db.SearchSource("Netflix")
	if err != nil {
		t.Errorf("SearchSource error: %s", err)
	}
	data, err := helper.ToJSON(source)
	if err != nil {
		t.Errorf("SearchSource - ToJSON - error: %s", err)
	}
	t.Log(data)
}
