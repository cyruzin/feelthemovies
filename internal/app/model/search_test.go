package model

import "testing"

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
	_, err = db.SearchUser("Cyro")

	if err != nil {
		t.Errorf("SearchUser error: %s", err)
	}
}

func TestSearchGenreModel(t *testing.T) {
	_, err = db.SearchGenre("Comedy")

	if err != nil {
		t.Errorf("SearchGenre error: %s", err)
	}
}

func TestSearchKeywordModel(t *testing.T) {
	_, err = db.SearchKeyword("Shark")

	if err != nil {
		t.Errorf("SearchKeyword error: %s", err)
	}
}

func TestSearchSourceModel(t *testing.T) {
	_, err = db.SearchSource("Netflix")

	if err != nil {
		t.Errorf("SearchSource error: %s", err)
	}
}
