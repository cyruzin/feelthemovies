package model

import (
	"testing"
	"time"
)

func TestGetRecommendationsModel(t *testing.T) {
	_, err := db.GetRecommendations(0, 10)
	if err != nil {
		t.Errorf("GetRecommendations error: %s", err)
	}
}

func TestGetRecommendationModel(t *testing.T) {
	_, err := db.GetRecommendation(1)
	if err != nil {
		t.Errorf("GetRecommendation error: %s", err)
	}
}

func TestCreateRecommendationModel(t *testing.T) {
	recItem := &Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := db.CreateRecommendation(recItem)
	if err != nil {
		t.Errorf("CreateRecommendationError error: %s", err)
	}
}

func TestUpdateRecommendationModel(t *testing.T) {
	recItem := &Recommendation{
		UserID:    1,
		Title:     "Aquaman",
		Type:      1,
		Body:      "The new Aquaman movie",
		Status:    2,
		Poster:    "ahs9qounasas",
		Backdrop:  "ajsopqwhasn",
		UpdatedAt: time.Now(),
	}
	_, err := db.UpdateRecommendation(1, recItem)
	if err != nil {
		t.Errorf("UpdateRecommendationItemError error: %s", err)
	}
}

func TestDeleteRecommendationModel(t *testing.T) {
	if err := db.DeleteRecommendation(3); err != nil {
		t.Errorf("DeleteRecommendation error: %s", err)
	}
}
