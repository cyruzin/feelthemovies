package model

import (
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

func TestGetRecommendationsModel(t *testing.T) {

	rec, err := db.GetRecommendations(0, 10)

	if err != nil {
		t.Errorf("GetRecommendations error: %s", err)
	}

	data, err := helper.ToJSONIndent(rec)

	if err != nil {
		t.Errorf("GetRecommendations error: %s", err)
	}

	t.Log(data)
}

func TestGetRecommendationModel(t *testing.T) {
	rec, err := db.GetRecommendation(1)

	if err != nil {
		t.Errorf("GetRecommendation error: %s", err)
	}

	data, err := helper.ToJSONIndent(rec)

	if err != nil {
		t.Errorf("GetRecommendation error: %s", err)
	}

	t.Log(data)
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

	ri, err := db.CreateRecommendation(recItem)

	if err != nil {
		t.Errorf("CreateRecommendationError error: %s", err)
	}

	data, err := helper.ToJSON(ri)

	if err != nil {
		t.Errorf("CreateRecommendationError - ToJSON error: %s", err)
	}

	t.Log(data)
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

	ri, err := db.UpdateRecommendation(1, recItem)

	if err != nil {
		t.Errorf("UpdateRecommendationItemError error: %s", err)
	}

	data, err := helper.ToJSON(ri)

	if err != nil {
		t.Errorf("UpdateRecommendationItemError - ToJSON error: %s", err)
	}

	t.Log(data)
}

func TestDeleteRecommendationModel(t *testing.T) {
	_, err = db.DeleteRecommendation(3)

	if err != nil {
		t.Errorf("DeleteRecommendation error: %s", err)
	}
}
