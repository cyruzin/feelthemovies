package model

import "testing"

func TestGetRecommendationsModel(t *testing.T) {

	_, err = db.GetRecommendations(0, 10)

	if err != nil {
		t.Errorf("GetRecommendations error: %s", err)
	}
}

func TestGetRecommendationModel(t *testing.T) {
	_, err = db.GetRecommendation(1)

	if err != nil {
		t.Errorf("GetRecommendation error: %s", err)
	}
}

func TestDeleteRecommendationModel(t *testing.T) {
	_, err = db.DeleteRecommendation(3)

	if err != nil {
		t.Errorf("DeleteRecommendation error: %s", err)
	}
}
